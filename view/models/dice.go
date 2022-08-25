package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dice struct {
	ID     primitive.ObjectID
	Number int
}
