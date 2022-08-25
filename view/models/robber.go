package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Robber struct {
	ID        primitive.ObjectID
	TerrainID primitive.ObjectID
	IsMoving  bool
}
