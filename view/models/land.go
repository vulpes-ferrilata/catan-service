package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Land struct {
	ID       primitive.ObjectID
	Q        int
	R        int
	Location string
}
