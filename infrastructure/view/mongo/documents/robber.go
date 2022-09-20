package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Robber struct {
	ID primitive.ObjectID `bson:"_id"`
}
