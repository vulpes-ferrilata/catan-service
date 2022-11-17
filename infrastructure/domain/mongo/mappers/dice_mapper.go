package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type diceMapper struct{}

func (d diceMapper) ToDocument(dice *models.Dice) (*documents.Dice, error) {
	if dice == nil {
		return nil, nil
	}

	return &documents.Dice{
		Document: documents.Document{
			ID: dice.GetID(),
		},
		Number: dice.GetNumber(),
	}, nil
}

func (d diceMapper) ToDomain(diceDocument *documents.Dice) (*models.Dice, error) {
	if diceDocument == nil {
		return nil, nil
	}

	dice := models.DiceBuilder{}.
		SetID(diceDocument.ID).
		SetNumber(diceDocument.Number).
		Create()

	return dice, nil
}
