package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type achievementMapper struct{}

func (a achievementMapper) ToDocument(achievement *models.Achievement) (*documents.Achievement, error) {
	if achievement == nil {
		return nil, nil
	}

	return &documents.Achievement{
		Document: documents.Document{
			ID: achievement.GetID(),
		},
		Type: achievement.GetType().String(),
	}, nil
}

func (a achievementMapper) ToDomain(achievementDocument *documents.Achievement) (*models.Achievement, error) {
	if achievementDocument == nil {
		return nil, nil
	}

	achievementType, err := models.NewAchievementType(achievementDocument.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievement := models.AchievementBuilder{}.
		SetID(achievementDocument.ID).
		SetType(achievementType).
		Create()

	return achievement, nil
}
