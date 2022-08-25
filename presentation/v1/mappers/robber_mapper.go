package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toRobberResponse(robber *models.Robber) *catan.RobberResponse {
	if robber == nil {
		return nil
	}

	return &catan.RobberResponse{
		ID:        robber.ID.Hex(),
		TerrainID: robber.TerrainID.Hex(),
		IsMoving:  robber.IsMoving,
	}
}
