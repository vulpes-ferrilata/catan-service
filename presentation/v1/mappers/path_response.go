package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type pathMapper struct{}

func (l pathMapper) ToResponse(path *models.Path) (*responses.Path, error) {
	if path == nil {
		return nil, nil
	}

	return &responses.Path{
		ID:       path.ID.Hex(),
		Q:        int32(path.Q),
		R:        int32(path.R),
		Location: path.Location,
	}, nil
}
