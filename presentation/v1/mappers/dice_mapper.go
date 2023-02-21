package mappers

import (
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type diceMapper struct{}

func (d diceMapper) ToResponse(dice *models.Dice) (*pb_models.Dice, error) {
	if dice == nil {
		return nil, nil
	}

	return &pb_models.Dice{
		ID:     dice.ID.Hex(),
		Number: int32(dice.Number),
	}, nil
}
