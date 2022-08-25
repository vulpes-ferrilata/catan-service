package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toDevelopmentCardResponse(developmentCard *models.DevelopmentCard) *catan.DevelopmentCardResponse {
	if developmentCard == nil {
		return nil
	}

	return &catan.DevelopmentCardResponse{
		ID:     developmentCard.ID.Hex(),
		Type:   developmentCard.Type,
		Status: developmentCard.Status,
	}
}
