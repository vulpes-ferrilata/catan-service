package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type gameMapper struct{}

func (g gameMapper) ToResponse(gameView *models.Game) (*responses.Game, error) {
	if gameView == nil {
		return nil, nil
	}

	return &responses.Game{
		ID:             gameView.ID.Hex(),
		PlayerQuantity: int32(gameView.PlayerQuantity),
		Status:         gameView.Status,
	}, nil
}
