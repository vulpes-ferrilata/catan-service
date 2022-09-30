package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toPathResponse(path *models.Path) *responses.Path {
	if path == nil {
		return nil
	}

	return &responses.Path{
		ID:       path.ID.Hex(),
		Q:        int32(path.Q),
		R:        int32(path.R),
		Location: path.Location,
	}
}
