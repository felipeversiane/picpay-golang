package router

import (
	"net/http"

	_ "github.com/felipeversiane/picpay-golang.git/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		UserRoutes(v1)
		OrderRoutes(v1)

	}

	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/docs/*any", swagger.WrapHandler(swaggerFiles.Handler))

}
