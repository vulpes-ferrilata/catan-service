package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConstructionBuilder interface {
	SetID(id primitive.ObjectID) ConstructionBuilder
	SetType(constructionType constructionType) ConstructionBuilder
	SetLand(land *Land) ConstructionBuilder
	Create() *Construction
}

func NewConstructionBuilder() ConstructionBuilder {
	return &constructionBuilder{}
}

type constructionBuilder struct {
	id               primitive.ObjectID
	constructionType constructionType
	land             *Land
}

func (a *constructionBuilder) SetID(id primitive.ObjectID) ConstructionBuilder {
	a.id = id

	return a
}

func (a *constructionBuilder) SetType(constructionType constructionType) ConstructionBuilder {
	a.constructionType = constructionType

	return a
}

func (a *constructionBuilder) SetLand(land *Land) ConstructionBuilder {
	a.land = land

	return a
}

func (a constructionBuilder) Create() *Construction {
	return &Construction{
		aggregate: aggregate{
			id: a.id,
		},
		constructionType: a.constructionType,
		land:             a.land,
	}
}
