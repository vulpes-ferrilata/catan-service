package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Road struct {
	ID   primitive.ObjectID `bson:"_id"`
	Path *Path              `bson:"path"`
}
