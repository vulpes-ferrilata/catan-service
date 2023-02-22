package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	view_models "github.com/vulpes-ferrilata/catan-service/view/models"
)

type GameOverviewMapper struct{}

func (g GameOverviewMapper) ToView(game *models.Game) (*view_models.GameOverview, error) {
	if game == nil {
		return nil, nil
	}

	allPlayers := append(game.GetPlayers(), game.GetActivePlayer())

	return &view_models.GameOverview{
		ID:             game.GetID(),
		PlayerQuantity: len(allPlayers),
		Status:         game.GetStatus().String(),
	}, nil
}
