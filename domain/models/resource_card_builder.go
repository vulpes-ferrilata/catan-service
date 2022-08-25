package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResourceCardBuilder interface {
	SetID(id primitive.ObjectID) ResourceCardBuilder
	SetType(resourceCardType resourceCardType) ResourceCardBuilder
	SetIsSelected(isSelected bool) ResourceCardBuilder
	Create() *ResourceCard
}

func NewResourceCardBuilder() ResourceCardBuilder {
	return &resourceCardBuilder{}
}

type resourceCardBuilder struct {
	id               primitive.ObjectID
	resourceCardType resourceCardType
	isSelected       bool
}

func (r *resourceCardBuilder) SetID(id primitive.ObjectID) ResourceCardBuilder {
	r.id = id

	return r
}

func (r *resourceCardBuilder) SetType(resourceCardType resourceCardType) ResourceCardBuilder {
	r.resourceCardType = resourceCardType

	return r
}

func (r *resourceCardBuilder) SetIsSelected(isSelected bool) ResourceCardBuilder {
	r.isSelected = isSelected

	return r
}

func (r resourceCardBuilder) Create() *ResourceCard {
	return &ResourceCard{
		aggregate: aggregate{
			id: r.id,
		},
		resourceCardType: r.resourceCardType,
		isSelected:       r.isSelected,
	}
}
