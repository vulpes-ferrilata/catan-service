package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type developmentCardMapper struct{}

func (d developmentCardMapper) ToResponse(developmentCard *models.DevelopmentCard) (*responses.DevelopmentCard, error) {
	if developmentCard == nil {
		return nil, nil
	}

	return &responses.DevelopmentCard{
		ID:     developmentCard.ID.Hex(),
		Type:   developmentCard.Type,
		Status: developmentCard.Status,
	}, nil
}
