package main

import (
	"log"

	"github.com/ManuelLecaro/erpcore/internal/infra/server"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()

	sugar.Info("Initializing app")

	app := server.NewApp()

	sugar.Info("App running")

	app.RunApp()
}
