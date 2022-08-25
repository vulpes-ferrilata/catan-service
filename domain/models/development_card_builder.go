package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DevelopmentCardBuilder interface {
	SetID(id primitive.ObjectID) DevelopmentCardBuilder
	SetType(developmentCardType developmentCardType) DevelopmentCardBuilder
	SetStatus(status developmentCardStatus) DevelopmentCardBuilder
	Create() *DevelopmentCard
}

func NewDevelopmentCardBuilder() DevelopmentCardBuilder {
	return &developmentCardBuilder{}
}

type developmentCardBuilder struct {
	id                  primitive.ObjectID
	developmentCardType developmentCardType
	status              developmentCardStatus
}

func (a *developmentCardBuilder) SetID(id primitive.ObjectID) DevelopmentCardBuilder {
	a.id = id

	return a
}

func (a *developmentCardBuilder) SetType(developmentCardType developmentCardType) DevelopmentCardBuilder {
	a.developmentCardType = developmentCardType

	return a
}

func (a *developmentCardBuilder) SetStatus(developmentCardStatus developmentCardStatus) DevelopmentCardBuilder {
	a.status = developmentCardStatus

	return a
}

func (a developmentCardBuilder) Create() *DevelopmentCard {
	return &DevelopmentCard{
		aggregate: aggregate{
			id: a.id,
		},
		developmentCardType: a.developmentCardType,
		status:              a.status,
	}
}
