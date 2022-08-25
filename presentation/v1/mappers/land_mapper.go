package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toLandResponse(land *models.Land) *catan.LandResponse {
	if land == nil {
		return nil
	}

	return &catan.LandResponse{
		ID:       land.ID.Hex(),
		Q:        int32(land.Q),
		R:        int32(land.R),
		Location: land.Location,
	}
}
