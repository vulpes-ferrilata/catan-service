package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

func toDiceDocument(dice *models.Dice) *documents.Dice {
	if dice == nil {
		return nil
	}

	return &documents.Dice{
		Document: documents.Document{
			ID: dice.GetID(),
		},
		Number: dice.GetNumber(),
	}
}

func toDiceDomain(diceDocument *documents.Dice) *models.Dice {
	if diceDocument == nil {
		return nil
	}

	dice := models.NewDiceBuilder().
		SetID(diceDocument.ID).
		SetNumber(diceDocument.Number).
		Create()

	return dice
}
