package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResourceCardBuilder struct {
	id               primitive.ObjectID
	resourceCardType ResourceCardType
	offering         bool
}

func (r ResourceCardBuilder) SetID(id primitive.ObjectID) ResourceCardBuilder {
	r.id = id

	return r
}

func (r ResourceCardBuilder) SetType(resourceCardType ResourceCardType) ResourceCardBuilder {
	r.resourceCardType = resourceCardType

	return r
}

func (r ResourceCardBuilder) SetOffering(offering bool) ResourceCardBuilder {
	r.offering = offering

	return r
}

func (r ResourceCardBuilder) Create() *ResourceCard {
	return &ResourceCard{
		aggregate: aggregate{
			id: r.id,
		},
		resourceCardType: r.resourceCardType,
		offering:         r.offering,
	}
}
