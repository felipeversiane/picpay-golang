package router

import (
	"github.com/felipeversiane/picpay-golang.git/config/db"
	"github.com/felipeversiane/picpay-golang.git/internal/handler"
	"github.com/felipeversiane/picpay-golang.git/internal/repository"
	"github.com/felipeversiane/picpay-golang.git/internal/service"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	repo := repository.NewUserRepository(db.Conn)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	user := r.Group("/user")
	{
		user.POST("/insert_user", handler.InsertUserHandler)
		user.GET("/find_user_by_document/:document", handler.FindUserByDocumentHandler)
		user.GET("/find_user_by_email/:email", handler.FindUserByEmailHandler)

	}

	return user
}
