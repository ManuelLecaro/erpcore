package main

import (
	"log"

	"github.com/ManuelLecaro/erpcore/internal/infra/prov"
	"github.com/ManuelLecaro/erpcore/internal/infra/server"
	"go.uber.org/zap"
)

const AppName = "erpcore"

func main() {
	container, err := prov.BuildContainer()
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(func(graphServer *server.GraphHandler, app *server.APIServer) {
		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}

		sugar := logger.Sugar()

		sugar.Info("Initializing app")

		graphServer.Init(AppName)

		sugar.Info("App running")

		app.RunApp()
	})

	if err != nil {
		log.Fatal(err)
	}
}
