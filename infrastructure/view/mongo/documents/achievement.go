package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Achievement struct {
	ID   primitive.ObjectID `bson:"_id"`
	Type string             `bson:"type"`
}
