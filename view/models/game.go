package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	ID             primitive.ObjectID
	PlayerQuantity int
	Status         string
}
