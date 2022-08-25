package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Construction struct {
	ID   primitive.ObjectID `bson:"_id"`
	Type string             `bson:"type"`
	Land *Land              `bson:"land"`
}
