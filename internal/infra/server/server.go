package server

// TODO: implement a DI library so it can be more readable and easier for the dependency injection steps.

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ManuelLecaro/erpcore/internal/infra/configuration"
	"github.com/ManuelLecaro/erpcore/internal/infra/graph"
	"github.com/gin-gonic/gin"
)

// APIServer is the structure that holds the infrastructure needed on the dependency injection phase.
type APIServer struct {
	Server       *gin.Engine
	StatusReady  chan bool
	StatusKilled chan bool
	Graph        *graph.Resolver
	ServerPort   string
	AppName      string
}

// NewApp initializes the app
func NewApp() *APIServer {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	a := &APIServer{
		Server:       server,
		StatusReady:  make(chan bool),
		StatusKilled: make(chan bool),
	}

	a.initInfrastructure()

	return a
}

func (a *APIServer) initInfrastructure() {
	configurations := configuration.LoadConfiguration()

	a.AppName = configurations.GetString("NAME")
	a.ServerPort = configurations.GetString("PORT")
}

func (a *APIServer) runServer() {
	log.Println("running app server at port: ", a.ServerPort)
	if err := a.Server.Run(fmt.Sprintf(":%s", a.ServerPort)); err != nil {
		log.Println("Shutting down the server")
	}
}

func (a *APIServer) KillApp() {
	if a.StatusKilled != nil {
		a.StatusKilled <- true
	}
}

// RunApp initializes the server and consumers
func (a *APIServer) RunApp() {
	shutDown := make(chan bool)
	go a.runServer()
	if a.StatusReady != nil {
		a.StatusReady <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	close(shutDown)
	a.KillApp()
}
