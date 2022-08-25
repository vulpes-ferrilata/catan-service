package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Robber struct {
	aggregate
	terrainID primitive.ObjectID
	isMoving  bool
}

func (r Robber) GetTerrainID() primitive.ObjectID {
	return r.terrainID
}

func (r Robber) IsMoving() bool {
	return r.isMoving
}
