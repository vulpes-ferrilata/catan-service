package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toResourceCardResponse(resourceCard *models.ResourceCard) *responses.ResourceCard {
	if resourceCard == nil {
		return nil
	}

	return &responses.ResourceCard{
		ID:         resourceCard.ID.Hex(),
		Type:       resourceCard.Type,
		IsSelected: resourceCard.IsSelected,
	}
}
