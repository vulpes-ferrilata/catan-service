package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toHarborResponse(harbor *models.Harbor) *catan.HarborResponse {
	if harbor == nil {
		return nil
	}

	return &catan.HarborResponse{
		ID:   harbor.ID.Hex(),
		Q:    int32(harbor.Q),
		R:    int32(harbor.R),
		Type: harbor.Type,
	}
}
