package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Path struct {
	ID       primitive.ObjectID `bson:"_id"`
	Q        int                `bson:"q"`
	R        int                `bson:"r"`
	Location string             `bson:"location"`
}
