package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toPathResponse(path *models.Path) *catan.PathResponse {
	if path == nil {
		return nil
	}

	return &catan.PathResponse{
		ID:       path.ID.Hex(),
		Q:        int32(path.Q),
		R:        int32(path.R),
		Location: path.Location,
	}
}
