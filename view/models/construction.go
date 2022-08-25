package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Construction struct {
	ID   primitive.ObjectID
	Type string
	Land *Land
}
