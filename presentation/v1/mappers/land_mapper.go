package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type landMapper struct{}

func (l landMapper) ToResponse(land *models.Land) (*responses.Land, error) {
	if land == nil {
		return nil, nil
	}

	return &responses.Land{
		ID:       land.ID.Hex(),
		Q:        int32(land.Q),
		R:        int32(land.R),
		Location: land.Location,
	}, nil
}
