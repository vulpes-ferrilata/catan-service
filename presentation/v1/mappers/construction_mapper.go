package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type constructionMapper struct{}

func (d constructionMapper) ToResponse(construction *models.Construction) (*responses.Construction, error) {
	if construction == nil {
		return nil, nil
	}

	land, err := landMapper{}.ToResponse(construction.Land)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &responses.Construction{
		ID:   construction.ID.Hex(),
		Type: construction.Type,
		Land: land,
	}, nil
}
