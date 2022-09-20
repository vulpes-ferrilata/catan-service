package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DevelopmentCardBuilder struct {
	id                  primitive.ObjectID
	developmentCardType DevelopmentCardType
	status              developmentCardStatus
}

func (a DevelopmentCardBuilder) SetID(id primitive.ObjectID) DevelopmentCardBuilder {
	a.id = id

	return a
}

func (a DevelopmentCardBuilder) SetType(developmentCardType DevelopmentCardType) DevelopmentCardBuilder {
	a.developmentCardType = developmentCardType

	return a
}

func (a DevelopmentCardBuilder) SetStatus(developmentCardStatus developmentCardStatus) DevelopmentCardBuilder {
	a.status = developmentCardStatus

	return a
}

func (a DevelopmentCardBuilder) Create() *DevelopmentCard {
	return &DevelopmentCard{
		aggregate: aggregate{
			id: a.id,
		},
		developmentCardType: a.developmentCardType,
		status:              a.status,
	}
}
