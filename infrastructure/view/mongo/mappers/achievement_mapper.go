package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type achievementMapper struct{}

func (a achievementMapper) ToView(achievementDocument *documents.Achievement) (*models.Achievement, error) {
	if achievementDocument == nil {
		return nil, nil
	}

	return &models.Achievement{
		ID:   achievementDocument.ID,
		Type: achievementDocument.Type,
	}, nil
}
