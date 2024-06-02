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
	FindUserByIDHandler(c *gin.Context)
	FindUserByEmailHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
	InsertUserHandler(c *gin.Context)
	UpdateUserHandler(c *gin.Context)
}

// FindUserByIDHandler retrieves user information based on the provided user ID.
// @Summary Find User by ID
// @Description Retrieves user details based on the user ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "ID of the user to be retrieved"
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} http_error.HttpError "Error: Invalid user ID"
// @Failure 404 {object} http_error.HttpError "User not found"
// @Router /user/{id} [get]
func (uh userHandler) FindUserByIDHandler(c *gin.Context) {
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

	user, err := uh.userService.FindUserByIDService(id, ctxTimeout)
	if err != nil {
		logger.Error("Error finding user by ID", err, zap.String("journey", "findUserByID"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// FindUserByDocumentHandler retrieves user information based on the provided user document.
// @Summary Find User by Document
// @Description Retrieves user details based on the user document provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param document path string true "Document of the user to be retrieved"
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 404 {object} http_error.HttpError "User not found"
// @Router /user/find_user_by_document/{document} [get]
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

// FindUserByEmailHandler retrieves user information based on the provided user email.
// @Summary Find User by Email
// @Description Retrieves user details based on the user email provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param email path string true "Email of the user to be retrieved"
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 404 {object} http_error.HttpError "User not found"
// @Router /user/find_user_by_email/{email} [get]
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

// DeleteUserHandler deletes a user with the specified ID.
// @Summary Delete User
// @Description Deletes a user based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "ID of the user to be deleted"
// @Success 200
// @Failure 400 {object} http_error.HttpError
// @Failure 500 {object} http_error.HttpError
// @Router /user/{id} [delete]
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

	c.JSON(http.StatusNoContent, gin.H{})
}

// InsertUserHandler Creates a new user
// @Summary Insert a new user
// @Description Insert a new user with the provided user information
// @Tags Users
// @Accept json
// @Produce json
// @Param userRequest body request.UserRequest true "User information for registration"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} http_error.HttpError
// @Failure 500 {object} http_error.HttpError
// @Router /user [post]
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
	c.JSON(http.StatusCreated, result)
}

// UpdateUserHandler updates user information with the specified ID.
// @Summary Update User
// @Description Updates user details based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "ID of the user to be updated"
// @Param userRequest body request.UserUpdateRequest true "User information for update"
// @Success 200
// @Failure 400 {object} http_error.HttpError
// @Failure 500 {object} http_error.HttpError
// @Router /user/{id} [put]
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

	result, serviceErr := uh.userService.UpdateUserService(id, domain, ctxTimeout)
	if serviceErr != nil {
		logger.Error(
			"Error trying to call updateUser service",
			serviceErr,
			zap.String("journey", "updateUser"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
