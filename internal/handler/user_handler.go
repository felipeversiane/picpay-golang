package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/felipeversiane/picpay-golang.git/config/logger"
	"github.com/felipeversiane/picpay-golang.git/config/validation"
	domain "github.com/felipeversiane/picpay-golang.git/internal"
	"github.com/felipeversiane/picpay-golang.git/internal/entity/request"
	"github.com/felipeversiane/picpay-golang.git/internal/service"
	"github.com/gin-gonic/gin"
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
	FindUserByIDHandler(c *gin.Context)
	FindUserByEmailHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
	InsertUserHandler(c *gin.Context)
	UpdateUserHandler(c *gin.Context)
}

func (uh userHandler) FindUserByIDHandler(c *gin.Context) {

}

func (uh userHandler) FindUserByEmailHandler(c *gin.Context) {

}

func (uh userHandler) DeleteUserHandler(c *gin.Context) {

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

}
