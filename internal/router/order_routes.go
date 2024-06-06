package router

import (
	"github.com/felipeversiane/picpay-golang.git/config/db"
	"github.com/felipeversiane/picpay-golang.git/internal/handler"
	"github.com/felipeversiane/picpay-golang.git/internal/repository"
	"github.com/felipeversiane/picpay-golang.git/internal/service"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	order_repo := repository.NewOrderRepository(db.Conn)
	user_repo := repository.NewUserRepository(db.Conn)
	user_service := service.NewUserService(user_repo)
	order_service := service.NewOrderService(order_repo, user_service)
	handler := handler.NewOrderHandler(order_service)

	order := r.Group("/order")
	{
		order.POST("/", handler.InsertOrderHandler)
		order.GET("/:id", handler.FindOrderByIDHandler)
	}

	return order
}
