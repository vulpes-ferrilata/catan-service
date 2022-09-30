package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toConstructionResponse(construction *models.Construction) *responses.Construction {
	if construction == nil {
		return nil
	}

	land := toLandResponse(construction.Land)

	return &responses.Construction{
		ID:   construction.ID.Hex(),
		Type: construction.Type,
		Land: land,
	}
}
