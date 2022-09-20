package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResourceCardBuilder struct {
	id               primitive.ObjectID
	resourceCardType ResourceCardType
	isSelected       bool
}

func (r ResourceCardBuilder) SetID(id primitive.ObjectID) ResourceCardBuilder {
	r.id = id

	return r
}

func (r ResourceCardBuilder) SetType(resourceCardType ResourceCardType) ResourceCardBuilder {
	r.resourceCardType = resourceCardType

	return r
}

func (r ResourceCardBuilder) SetIsSelected(isSelected bool) ResourceCardBuilder {
	r.isSelected = isSelected

	return r
}

func (r ResourceCardBuilder) Create() *ResourceCard {
	return &ResourceCard{
		aggregate: aggregate{
			id: r.id,
		},
		resourceCardType: r.resourceCardType,
		isSelected:       r.isSelected,
	}
}
