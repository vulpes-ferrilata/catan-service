package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type diceMapper struct{}

func (d diceMapper) ToResponse(dice *models.Dice) (*responses.Dice, error) {
	if dice == nil {
		return nil, nil
	}

	return &responses.Dice{
		ID:     dice.ID.Hex(),
		Number: int32(dice.Number),
	}, nil
}
