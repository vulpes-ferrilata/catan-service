package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AchievementBuilder struct {
	id              primitive.ObjectID
	achievementType AchievementType
}

func (a AchievementBuilder) SetID(id primitive.ObjectID) AchievementBuilder {
	a.id = id

	return a
}

func (a AchievementBuilder) SetType(achievementType AchievementType) AchievementBuilder {
	a.achievementType = achievementType

	return a
}

func (a AchievementBuilder) Create() *Achievement {
	return &Achievement{
		aggregate: aggregate{
			id: a.id,
		},
		achievementType: a.achievementType,
	}
}
