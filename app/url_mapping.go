package app

import (
	"github.com/helyus1412/messaging-api/config"
	"github.com/helyus1412/messaging-api/handler"
	repo "github.com/helyus1412/messaging-api/repository"
	"github.com/helyus1412/messaging-api/services"
)

var (
	db         = config.Init_DB()
	repository = repo.NewRepository(db)
	service    = services.NewServices(repository)
	handlerAPI = handler.NewHandler(service)
)

func mapUrls() {
	v1 := router.Group("/api/v1/")
	{
		message := v1.Group("/messages")
		{
			message.GET("/", handlerAPI.GetAllMessages)
		}
	}
}
