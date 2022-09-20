package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConstructionBuilder struct {
	id               primitive.ObjectID
	constructionType ConstructionType
	land             *Land
}

func (a ConstructionBuilder) SetID(id primitive.ObjectID) ConstructionBuilder {
	a.id = id

	return a
}

func (a ConstructionBuilder) SetType(constructionType ConstructionType) ConstructionBuilder {
	a.constructionType = constructionType

	return a
}

func (a ConstructionBuilder) SetLand(land *Land) ConstructionBuilder {
	a.land = land

	return a
}

func (a ConstructionBuilder) Create() *Construction {
	return &Construction{
		aggregate: aggregate{
			id: a.id,
		},
		constructionType: a.constructionType,
		land:             a.land,
	}
}
