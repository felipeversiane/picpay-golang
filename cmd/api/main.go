package main

import (
	"context"
	"net/http"
	"os"

	"github.com/felipeversiane/picpay-golang.git/bootstrap"
	"github.com/felipeversiane/picpay-golang.git/config/db"
	"github.com/felipeversiane/picpay-golang.git/config/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	POSTGRES_URL = "POSTGRES_URL"
)

func init() {
	logger.Info("Initialize methods is going to run...",
		zap.String("journey", "Initialize"))

	err := bootstrap.Initialize()

	if err != nil {
		logger.Fatal("Initialize error: ", err,
			zap.String("journey", "Initialize"))
	}

	logger.Info("Initialize methods runned...",
		zap.String("journey", "Initialize"))
}

func main() {
	var err error
	ctx := context.Background()
	connectionString := os.Getenv(POSTGRES_URL)

	conn, err := db.NewConnection(ctx, connectionString)
	if err != nil {
		logger.Fatal("Database error: ", err,
			zap.String("journey", "Database Connection"))
	}
	defer conn.Close()

	logger.Info("Database connection completed",
		zap.String("journey", "Database Connection"))

	g := gin.New()
	g.Use(gin.Recovery())
	g.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	logger.Info("Routes initialized sucessfully.",
		zap.String("journey", "Initialize Routes"))
	g.Run()

}
