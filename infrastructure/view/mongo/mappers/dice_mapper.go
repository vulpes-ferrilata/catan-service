package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toDiceView(diceDocument *documents.Dice) *models.Dice {
	if diceDocument == nil {
		return nil
	}

	return &models.Dice{
		ID:     diceDocument.ID,
		Number: diceDocument.Number,
	}
}
