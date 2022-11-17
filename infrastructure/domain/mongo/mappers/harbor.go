package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type harborMapper struct{}

func (h harborMapper) ToDocument(harbor *models.Harbor) (*documents.Harbor, error) {
	if harbor == nil {
		return nil, nil
	}

	return &documents.Harbor{
		Document: documents.Document{
			ID: harbor.GetID(),
		},
		Q:    harbor.GetHex().GetQ(),
		R:    harbor.GetHex().GetR(),
		Type: harbor.GetType().String(),
	}, nil
}

func (h harborMapper) ToDomain(harborDocument *documents.Harbor) (*models.Harbor, error) {
	if harborDocument == nil {
		return nil, nil
	}

	hex := models.NewHex(harborDocument.Q, harborDocument.R)

	harborType, err := models.NewHarborType(harborDocument.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	harbor := models.HarborBuilder{}.
		SetID(harborDocument.ID).
		SetHex(hex).
		SetType(harborType).
		Create()

	return harbor, nil
}
