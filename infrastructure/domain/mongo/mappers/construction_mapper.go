package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type constructionMapper struct{}

func (c constructionMapper) ToDocument(construction *models.Construction) (*documents.Construction, error) {
	if construction == nil {
		return nil, nil
	}

	landDocument, err := landMapper{}.ToDocument(construction.GetLand())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &documents.Construction{
		Document: documents.Document{
			ID: construction.GetID(),
		},
		Type: construction.GetType().String(),
		Land: landDocument,
	}, nil
}

func (c constructionMapper) ToDomain(constructionDocument *documents.Construction) (*models.Construction, error) {
	if constructionDocument == nil {
		return nil, nil
	}

	constructionType, err := models.NewConstructionType(constructionDocument.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	land, err := landMapper{}.ToDomain(constructionDocument.Land)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	construction := models.ConstructionBuilder{}.
		SetID(constructionDocument.ID).
		SetType(constructionType).
		SetLand(land).
		Create()

	return construction, nil
}
