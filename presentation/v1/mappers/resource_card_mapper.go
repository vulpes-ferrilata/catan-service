package mappers

import (
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type resourceCardMapper struct{}

func (r resourceCardMapper) ToResponse(resourceCard *models.ResourceCard) (*pb_models.ResourceCard, error) {
	if resourceCard == nil {
		return nil, nil
	}

	return &pb_models.ResourceCard{
		ID:       resourceCard.ID.Hex(),
		Type:     resourceCard.Type,
		Offering: resourceCard.Offering,
	}, nil
}
