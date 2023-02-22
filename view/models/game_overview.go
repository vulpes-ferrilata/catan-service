package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GameOverview struct {
	ID             primitive.ObjectID
	PlayerQuantity int
	Status         string
}
