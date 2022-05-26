package results

import (
	"github.com/VulpesFerrilata/catan-service/infrastructure/bus"
	"go.uber.org/dig"
)

type CommandHandlerResult struct {
	dig.Out

	CommandHandler bus.CommandHandler `group:"commandBus"`
}
