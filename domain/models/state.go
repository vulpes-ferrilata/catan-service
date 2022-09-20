package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type state interface {
	newPlayer(userID primitive.ObjectID) error
	startGame(userID primitive.ObjectID) error
	phase
	playKnightCard(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error
	playRoadBuildingCard(userID primitive.ObjectID, pathIDs []primitive.ObjectID) error
	playYearOfPlentyCard(userID primitive.ObjectID, resourceCardTypes []ResourceCardType) error
	playMonopolyCard(userID primitive.ObjectID, resourceCardType ResourceCardType) error
}
