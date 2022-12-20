package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type state interface {
	newPlayer(userID primitive.ObjectID) error
	startGame(userID primitive.ObjectID) error
	phase
	playKnightCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error
	playRoadBuildingCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, pathIDs []primitive.ObjectID) error
	playYearOfPlentyCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, resourceCardTypes []ResourceCardType) error
	playMonopolyCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, resourceCardType ResourceCardType) error
	playVictoryPointCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID) error
}
