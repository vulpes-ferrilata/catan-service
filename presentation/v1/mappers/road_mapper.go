package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type roadMapper struct{}

func (r roadMapper) ToResponse(road *models.Road) (*responses.Road, error) {
	if road == nil {
		return nil, nil
	}

	path, err := pathMapper{}.ToResponse(road.Path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &responses.Road{
		ID:   road.ID.Hex(),
		Path: path,
	}, nil
}
