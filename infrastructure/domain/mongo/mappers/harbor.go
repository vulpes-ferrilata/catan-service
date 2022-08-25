package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

func toHarborDocument(harbor *models.Harbor) *documents.Harbor {
	if harbor == nil {
		return nil
	}

	return &documents.Harbor{
		Document: documents.Document{
			ID: harbor.GetID(),
		},
		TerrainID: harbor.GetTerrainID(),
		Q:         harbor.GetHex().GetQ(),
		R:         harbor.GetHex().GetR(),
		Type:      string(harbor.GetType()),
	}
}

func toHarborDomain(harborDocument *documents.Harbor) (*models.Harbor, error) {
	if harborDocument == nil {
		return nil, nil
	}

	hex := models.NewHex(harborDocument.Q, harborDocument.R)

	harborType, err := models.NewHarborType(harborDocument.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	harbor := models.NewHarborBuilder().
		SetID(harborDocument.ID).
		SetTerrainID(harborDocument.TerrainID).
		SetHex(hex).
		SetHarborType(harborType).
		Create()

	return harbor, nil
}
