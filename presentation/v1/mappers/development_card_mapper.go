package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toDevelopmentCardResponse(developmentCard *models.DevelopmentCard) *responses.DevelopmentCard {
	if developmentCard == nil {
		return nil
	}

	return &responses.DevelopmentCard{
		ID:     developmentCard.ID.Hex(),
		Type:   developmentCard.Type,
		Status: developmentCard.Status,
	}
}
