package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type resourceCardMapper struct{}

func (r resourceCardMapper) ToResponse(resourceCard *models.ResourceCard) (*responses.ResourceCard, error) {
	if resourceCard == nil {
		return nil, nil
	}

	return &responses.ResourceCard{
		ID:       resourceCard.ID.Hex(),
		Type:     resourceCard.Type,
		Offering: resourceCard.Offering,
	}, nil
}
