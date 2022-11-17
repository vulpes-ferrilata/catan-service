package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	ID             primitive.ObjectID `bson:"_id"`
	PlayerQuantity int                `bson:"player_quantity"`
	Status         string             `bson:"status"`
}
