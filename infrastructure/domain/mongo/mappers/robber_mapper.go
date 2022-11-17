package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type robberMapper struct{}

func (r robberMapper) ToDocument(robber *models.Robber) (*documents.Robber, error) {
	if robber == nil {
		return nil, nil
	}

	return &documents.Robber{
		Document: documents.Document{
			ID: robber.GetID(),
		},
	}, nil
}

func (r robberMapper) ToDomain(robberDocument *documents.Robber) (*models.Robber, error) {
	if robberDocument == nil {
		return nil, nil
	}

	robber := models.RobberBuilder{}.
		SetID(robberDocument.ID).
		Create()

	return robber, nil
}
