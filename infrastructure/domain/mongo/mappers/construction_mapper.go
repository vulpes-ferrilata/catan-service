package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

func toConstructionDocument(construction *models.Construction) *documents.Construction {
	if construction == nil {
		return nil
	}

	landDocument := toLandDocument(construction.GetLand())

	return &documents.Construction{
		Document: documents.Document{
			ID: construction.GetID(),
		},
		Type: string(construction.GetType()),
		Land: landDocument,
	}
}

func toConstructionDomain(constructionDocument *documents.Construction) (*models.Construction, error) {
	if constructionDocument == nil {
		return nil, nil
	}

	constructionType, err := models.NewConstructionType(constructionDocument.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	land, err := toLandDomain(constructionDocument.Land)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	construction := models.NewConstructionBuilder().
		SetID(constructionDocument.ID).
		SetType(constructionType).
		SetLand(land).
		Create()

	return construction, nil
}
