package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResourceCard struct {
	ID         primitive.ObjectID `bson:"_id"`
	Type       string             `bson:"type"`
	IsSelected bool               `bson:"is_selected"`
}
