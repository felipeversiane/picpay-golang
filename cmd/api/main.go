package main

import (
	"github.com/felipeversiane/picpay-golang.git/bootstrap"
	"github.com/felipeversiane/picpay-golang.git/config/logger"
	"go.uber.org/zap"
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

}
