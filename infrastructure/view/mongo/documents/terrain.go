package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Terrain struct {
	ID     primitive.ObjectID `bson:"_id"`
	Q      int                `bson:"q"`
	R      int                `bson:"r"`
	Number int
	Type   string
}
