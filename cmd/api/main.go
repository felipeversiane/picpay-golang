package main

import (
	"context"
	"os"

	"github.com/felipeversiane/picpay-golang.git/config/db"
	"github.com/felipeversiane/picpay-golang.git/config/logger"
	"github.com/felipeversiane/picpay-golang.git/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var (
	POSTGRES_URL = "POSTGRES_URL"
)

func init() {
	logger.Info("Initialize methods is going to run...",
		zap.String("journey", "Initialize"))

	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			logger.Fatal("Error loading .env file: ", err,
				zap.String("journey", "Loading .env"))
		}
	}

	logger.Info("Initialize methods runned...",
		zap.String("journey", "Initialize"))
}

// @title PicPay Challange
// @version 1.0
// @description REST API for a PicPay Challange.
// @host picpay-golang.onrender.com/docs/index.html
// @BasePath /api/v1
// @schemes http
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email felipeversiane09@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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
	router.InitRoutes(g)
	logger.Info("Routes initialized sucessfully.",
		zap.String("journey", "Initialize Routes"))
	g.Run()

}
