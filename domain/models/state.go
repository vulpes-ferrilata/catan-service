package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type state interface {
	newPlayer(userID primitive.ObjectID) error
	startGame(userID primitive.ObjectID) error
	phase
}
