package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Harbor struct {
	ID   primitive.ObjectID `bson:"_id"`
	Q    int                `bson:"q"`
	R    int                `bson:"r"`
	Type string             `bson:"type"`
}
