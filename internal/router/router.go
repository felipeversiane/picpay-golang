package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		UserRoutes(v1)
	}

	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

}
