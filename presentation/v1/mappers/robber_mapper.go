package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type robberMapper struct{}

func (r robberMapper) ToResponse(robber *models.Robber) (*responses.Robber, error) {
	if robber == nil {
		return nil, nil
	}

	return &responses.Robber{
		ID: robber.ID.Hex(),
	}, nil
}
