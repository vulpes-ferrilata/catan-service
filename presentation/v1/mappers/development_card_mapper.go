package mappers

import (
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type developmentCardMapper struct{}

func (d developmentCardMapper) ToResponse(developmentCard *models.DevelopmentCard) (*pb_models.DevelopmentCard, error) {
	if developmentCard == nil {
		return nil, nil
	}

	return &pb_models.DevelopmentCard{
		ID:     developmentCard.ID.Hex(),
		Type:   developmentCard.Type,
		Status: developmentCard.Status,
	}, nil
}
