package rest

import (
	"github.com/VulpesFerrilata/catan-service/infrastructure"
	"github.com/VulpesFerrilata/catan-service/presentation/rest/controllers"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := infrastructure.NewContainer()

	//Presentation layer
	container.Provide(controllers.NewCatanController)
	container.Provide(controllers.NewCatanWebsocketController)

	container.Provide(NewRouter)
	container.Provide(NewApp)
	container.Provide(NewServer)
	container.Provide(NewClient)
	container.Provide(NewService)

	return container
}
