package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type GameOverviewMapper struct{}

func (g GameOverviewMapper) ToDocument(gameOverview *models.GameOverview) (*documents.GameOverview, error) {
	if gameOverview == nil {
		return nil, nil
	}

	return &documents.GameOverview{
		ID:             gameOverview.ID,
		PlayerQuantity: gameOverview.PlayerQuantity,
		Status:         gameOverview.Status,
	}, nil
}

func (g GameOverviewMapper) ToView(gameDocument *documents.Game) (*models.Game, error) {
	if gameDocument == nil {
		return nil, nil
	}

	return &models.Game{
		ID:             gameDocument.ID,
		PlayerQuantity: gameDocument.PlayerQuantity,
		Status:         gameDocument.Status,
	}, nil
}
