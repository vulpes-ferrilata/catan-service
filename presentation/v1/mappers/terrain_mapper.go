package mappers

import (
	"github.com/pkg/errors"
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type terrainMapper struct{}

func (t terrainMapper) ToResponse(terrain *models.Terrain) (*pb_models.Terrain, error) {
	if terrain == nil {
		return nil, nil
	}

	terrainResponse, err := harborMapper{}.ToResponse(terrain.Harbor)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	robberResponse, err := robberMapper{}.ToResponse(terrain.Robber)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb_models.Terrain{
		ID:     terrain.ID.Hex(),
		Q:      int32(terrain.Q),
		R:      int32(terrain.R),
		Number: int32(terrain.Number),
		Type:   terrain.Type,
		Harbor: terrainResponse,
		Robber: robberResponse,
	}, nil
}
