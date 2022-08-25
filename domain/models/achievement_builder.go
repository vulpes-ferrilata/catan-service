package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AchievementBuilder interface {
	SetID(id primitive.ObjectID) AchievementBuilder
	SetType(achievementType achievementType) AchievementBuilder
	Create() *Achievement
}

func NewAchievementBuilder() AchievementBuilder {
	return &achievementBuilder{}
}

type achievementBuilder struct {
	id              primitive.ObjectID
	achievementType achievementType
}

func (a *achievementBuilder) SetID(id primitive.ObjectID) AchievementBuilder {
	a.id = id

	return a
}

func (a *achievementBuilder) SetType(achievementType achievementType) AchievementBuilder {
	a.achievementType = achievementType

	return a
}

func (a achievementBuilder) Create() *Achievement {
	return &Achievement{
		aggregate: aggregate{
			id: a.id,
		},
		achievementType: a.achievementType,
	}
}
