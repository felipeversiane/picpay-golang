package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/felipeversiane/picpay-golang.git/config/http_error"
	"github.com/felipeversiane/picpay-golang.git/config/logger"
	"github.com/felipeversiane/picpay-golang.git/config/validation"
	domain "github.com/felipeversiane/picpay-golang.git/internal"
	"github.com/felipeversiane/picpay-golang.git/internal/entity/request"
	"github.com/felipeversiane/picpay-golang.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(
	userService service.UserService,
) UserHandler {
	return &userHandler{
		userService,
	}
}

type UserHandler interface {
	FindUserByDocumentHandler(c *gin.Context)
	FindUserByEmailHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
	InsertUserHandler(c *gin.Context)
	UpdateUserHandler(c *gin.Context)
}

func (uh userHandler) FindUserByDocumentHandler(c *gin.Context) {
	document := c.Param("document")

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user, err := uh.userService.FindUserByDocumentService(document, ctxTimeout)
	if err != nil {
		logger.Error("Error finding user by Document", err, zap.String("journey", "findUserByDocument"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh userHandler) FindUserByEmailHandler(c *gin.Context) {
	email := c.Param("email")

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user, err := uh.userService.FindUserByEmailService(email, ctxTimeout)
	if err != nil {
		logger.Error("Error finding user by Email", err, zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh userHandler) DeleteUserHandler(c *gin.Context) {
	id, parseError := uuid.Parse(c.Param("id"))
	if parseError != nil {
		logger.Error("Error trying to validate userId",
			parseError,
			zap.String("journey", "DeleteUser"),
		)
		errorMessage := http_error.NewBadRequestError(
			"The ID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	serviceError := uh.userService.DeleteUserService(id, ctxTimeout)
	if serviceError != nil {
		logger.Error("Error trying to call deleteUser service", serviceError, zap.String("journey", "DeleteUser"))
		c.JSON(serviceError.Code, serviceError)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (uh userHandler) InsertUserHandler(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := domain.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.IsMerchant,
		userRequest.Document,
		userRequest.Balance,
	)

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err := uh.userService.InsertUserService(ctxTimeout, domain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (uh userHandler) UpdateUserHandler(c *gin.Context) {
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("journey", "Update"),
		)
		errorMessage := http_error.NewBadRequestError(
			"The ID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	domain := domain.NewUserUpdateDomain(
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Balance,
		userRequest.IsMerchant,
	)

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := uh.userService.UpdateUserService(id, domain, ctxTimeout); err != nil {
		logger.Error(
			"Error trying to call updateUser service",
			err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
