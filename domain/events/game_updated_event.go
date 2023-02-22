package events

import "github.com/vulpes-ferrilata/catan-service/domain/models"

type GameUpdatedEvent struct {
	Game *models.Game
}
