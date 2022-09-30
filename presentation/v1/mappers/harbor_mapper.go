package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toHarborResponse(harbor *models.Harbor) *responses.Harbor {
	if harbor == nil {
		return nil
	}

	return &responses.Harbor{
		ID:   harbor.ID.Hex(),
		Q:    int32(harbor.Q),
		R:    int32(harbor.R),
		Type: harbor.Type,
	}
}
