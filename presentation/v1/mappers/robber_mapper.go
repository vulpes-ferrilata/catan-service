package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toRobberResponse(robber *models.Robber) *responses.Robber {
	if robber == nil {
		return nil
	}

	return &responses.Robber{
		ID: robber.ID.Hex(),
	}
}
