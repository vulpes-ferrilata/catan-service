package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toResourceCardResponse(resourceCard *models.ResourceCard) *catan.ResourceCardResponse {
	if resourceCard == nil {
		return nil
	}

	return &catan.ResourceCardResponse{
		ID:         resourceCard.ID.Hex(),
		Type:       resourceCard.Type,
		IsSelected: resourceCard.IsSelected,
	}
}
