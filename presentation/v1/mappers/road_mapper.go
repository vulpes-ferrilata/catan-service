package mappers

import (
	"github.com/pkg/errors"
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type roadMapper struct{}

func (r roadMapper) ToResponse(road *models.Road) (*pb_models.Road, error) {
	if road == nil {
		return nil, nil
	}

	path, err := pathMapper{}.ToResponse(road.Path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb_models.Road{
		ID:   road.ID.Hex(),
		Path: path,
	}, nil
}
