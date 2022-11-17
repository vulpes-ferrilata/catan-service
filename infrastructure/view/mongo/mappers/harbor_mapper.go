package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type harborMapper struct{}

func (h harborMapper) ToView(harborDocument *documents.Harbor) (*models.Harbor, error) {
	if harborDocument == nil {
		return nil, nil
	}

	return &models.Harbor{
		ID:   harborDocument.ID,
		Q:    harborDocument.Q,
		R:    harborDocument.R,
		Type: harborDocument.Type,
	}, nil
}
