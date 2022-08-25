package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Path struct {
	ID       primitive.ObjectID
	Q        int
	R        int
	Location string
}
