package mappers

import (
	"github.com/pkg/errors"
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type constructionMapper struct{}

func (d constructionMapper) ToResponse(construction *models.Construction) (*pb_models.Construction, error) {
	if construction == nil {
		return nil, nil
	}

	land, err := landMapper{}.ToResponse(construction.Land)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb_models.Construction{
		ID:   construction.ID.Hex(),
		Type: construction.Type,
		Land: land,
	}, nil
}
