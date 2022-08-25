package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toDiceResponse(dice *models.Dice) *catan.DiceResponse {
	if dice == nil {
		return nil
	}

	return &catan.DiceResponse{
		ID:     dice.ID.Hex(),
		Number: int32(dice.Number),
	}
}
