package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Road struct {
	ID   primitive.ObjectID
	Path *Path
}
