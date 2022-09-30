package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toDiceResponse(dice *models.Dice) *responses.Dice {
	if dice == nil {
		return nil
	}

	return &responses.Dice{
		ID:     dice.ID.Hex(),
		Number: int32(dice.Number),
	}
}
