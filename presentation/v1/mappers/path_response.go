package mappers

import (
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type pathMapper struct{}

func (l pathMapper) ToResponse(path *models.Path) (*pb_models.Path, error) {
	if path == nil {
		return nil, nil
	}

	return &pb_models.Path{
		ID:       path.ID.Hex(),
		Q:        int32(path.Q),
		R:        int32(path.R),
		Location: path.Location,
	}, nil
}
