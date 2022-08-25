package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RobberBuilder interface {
	SetID(id primitive.ObjectID) RobberBuilder
	SetTerrainID(terrainID primitive.ObjectID) RobberBuilder
	SetIsMoving(isMoving bool) RobberBuilder
	Create() *Robber
}

func NewRobberBuilder() RobberBuilder {
	return &robberBuilder{}
}

type robberBuilder struct {
	id        primitive.ObjectID
	terrainID primitive.ObjectID
	isMoving  bool
}

func (r *robberBuilder) SetID(id primitive.ObjectID) RobberBuilder {
	r.id = id

	return r
}

func (r *robberBuilder) SetTerrainID(terrainID primitive.ObjectID) RobberBuilder {
	r.terrainID = terrainID

	return r
}

func (r *robberBuilder) SetIsMoving(isMoving bool) RobberBuilder {
	r.isMoving = isMoving

	return r
}

func (r robberBuilder) Create() *Robber {
	return &Robber{
		aggregate: aggregate{
			id: r.id,
		},
		terrainID: r.terrainID,
		isMoving:  r.isMoving,
	}
}
