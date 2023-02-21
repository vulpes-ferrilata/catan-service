package mappers

import (
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type gameMapper struct{}

func (g gameMapper) ToResponse(gameView *models.Game) (*pb_models.Game, error) {
	if gameView == nil {
		return nil, nil
	}

	return &pb_models.Game{
		ID:             gameView.ID.Hex(),
		PlayerQuantity: int32(gameView.PlayerQuantity),
		Status:         gameView.Status,
	}, nil
}
