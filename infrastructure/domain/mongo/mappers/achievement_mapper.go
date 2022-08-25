package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

func toAchievementDocument(achievement *models.Achievement) *documents.Achievement {
	if achievement == nil {
		return nil
	}

	return &documents.Achievement{
		Document: documents.Document{
			ID: achievement.GetID(),
		},
		Type: string(achievement.GetType()),
	}
}

func toAchievementDomain(achievementDocument *documents.Achievement) (*models.Achievement, error) {
	if achievementDocument == nil {
		return nil, nil
	}

	achievementType, err := models.NewAchievementType(achievementDocument.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievement := models.NewAchievementBuilder().
		SetID(achievementDocument.ID).
		SetType(achievementType).
		Create()

	return achievement, nil
}
