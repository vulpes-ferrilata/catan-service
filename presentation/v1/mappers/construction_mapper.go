package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toConstructionResponse(construction *models.Construction) *catan.ConstructionResponse {
	if construction == nil {
		return nil
	}

	land := toLandResponse(construction.Land)

	return &catan.ConstructionResponse{
		ID:   construction.ID.Hex(),
		Type: construction.Type,
		Land: land,
	}
}
