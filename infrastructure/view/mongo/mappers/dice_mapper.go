package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type diceMapper struct{}

func (d diceMapper) ToView(diceDocument *documents.Dice) (*models.Dice, error) {
	if diceDocument == nil {
		return nil, nil
	}

	return &models.Dice{
		ID:     diceDocument.ID,
		Number: diceDocument.Number,
	}, nil
}
