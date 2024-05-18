package main

import (
	"context"
	"os"

	"github.com/felipeversiane/picpay-golang.git/bootstrap"
	"github.com/felipeversiane/picpay-golang.git/config/db"
	"github.com/felipeversiane/picpay-golang.git/config/logger"
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

}
