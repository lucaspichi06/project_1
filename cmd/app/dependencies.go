package app

import (
	"github.com/gin-gonic/gin"
	"github.com/project_1/cmd/controller"
	"github.com/project_1/cmd/infrastructure/db"
	"github.com/project_1/cmd/service"
)

type produceController interface {
	Add(*gin.Context)
	Fetch(*gin.Context)
	Delete(*gin.Context)
}

type statusController interface {
	HandlePing(*gin.Context)
}

type controllers struct {
	produce produceController
	status  statusController
}

func load() *controllers {
	// DB layer
	produceDB := db.NewDataBase()

	// Application layer
	produceService := service.NewProduceService(&produceDB)

	// UI layer
	produceController := controller.NewProduceController(&produceService)
	statusController := controller.NewStatusController()

	controllers := &controllers{
		produce: &produceController,
		status:  statusController,
	}

	return controllers
}
