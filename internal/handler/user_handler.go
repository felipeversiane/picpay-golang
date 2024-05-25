package handler

import (
	"github.com/felipeversiane/picpay-golang.git/internal/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

type UserHandler interface {
	FindUserByIDHandler(c *gin.Context)
	FindUserByEmail(c *gin.Context)
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

}

func (uh userHandler) UpdateUserHandler(c *gin.Context) {

}
