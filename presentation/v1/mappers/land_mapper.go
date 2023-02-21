package mappers

import (
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type landMapper struct{}

func (l landMapper) ToResponse(land *models.Land) (*pb_models.Land, error) {
	if land == nil {
		return nil, nil
	}

	return &pb_models.Land{
		ID:       land.ID.Hex(),
		Q:        int32(land.Q),
		R:        int32(land.R),
		Location: land.Location,
	}, nil
}
