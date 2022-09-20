package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Terrain struct {
	ID     primitive.ObjectID
	Q      int
	R      int
	Number int
	Type   string
	Harbor *Harbor
	Robber *Robber
}
