package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DevelopmentCard struct {
	ID     primitive.ObjectID
	Type   string
	Status string
}
