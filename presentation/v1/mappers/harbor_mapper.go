package mappers

import (
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type harborMapper struct{}

func (h harborMapper) ToResponse(harbor *models.Harbor) (*pb_models.Harbor, error) {
	if harbor == nil {
		return nil, nil
	}

	return &pb_models.Harbor{
		ID:   harbor.ID.Hex(),
		Q:    int32(harbor.Q),
		R:    int32(harbor.R),
		Type: harbor.Type,
	}, nil
}
