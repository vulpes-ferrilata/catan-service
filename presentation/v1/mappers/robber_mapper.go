package mappers

import (
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type robberMapper struct{}

func (r robberMapper) ToResponse(robber *models.Robber) (*pb_models.Robber, error) {
	if robber == nil {
		return nil, nil
	}

	return &pb_models.Robber{
		ID: robber.ID.Hex(),
	}, nil
}
