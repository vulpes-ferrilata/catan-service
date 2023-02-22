package events

import "github.com/vulpes-ferrilata/catan-service/domain/models"

type GameCreatedEvent struct {
	Game *models.Game
}
