package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type gameMapper struct{}

func (g gameMapper) ToView(gameDocument *documents.Game) (*models.Game, error) {
	if gameDocument == nil {
		return nil, nil
	}

	return &models.Game{
		ID:             gameDocument.ID,
		PlayerQuantity: gameDocument.PlayerQuantity,
		Status:         gameDocument.Status,
	}, nil
}
