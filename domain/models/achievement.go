package models

import (
	"github.com/VulpesFerrilata/catan-service/domain/models/common"
	"github.com/google/uuid"
)

func NewAchievement(id uuid.UUID, achievementType AchievementType) *Achievement {
	return &Achievement{
		Entity:          common.NewEntity(id),
		achievementType: achievementType,
	}
}

type Achievement struct {
	common.Entity
	achievementType AchievementType
}

func (a Achievement) GetType() AchievementType {
	return a.achievementType
}
