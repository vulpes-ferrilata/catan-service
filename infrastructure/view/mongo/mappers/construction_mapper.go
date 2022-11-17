package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type constructionMapper struct{}

func (c constructionMapper) ToView(constructionDocument *documents.Construction) (*models.Construction, error) {
	if constructionDocument == nil {
		return nil, nil
	}

	land, err := landMapper{}.ToView(constructionDocument.Land)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.Construction{
		ID:   constructionDocument.ID,
		Type: constructionDocument.Type,
		Land: land,
	}, nil
}
