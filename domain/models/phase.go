package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type phase interface {
	buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error
	rollDices(userID primitive.ObjectID) error
	moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error
	endTurn(userID primitive.ObjectID) error
	// buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error
	// buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error
	// 	upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error
	// 	buyDevelopmentCard(userID primitive.ObjectID) error
	// 	maritimeTrade(userID primitive.ObjectID, fromResourceCardType ResourceCardType, toResourceCardType ResourceCardType) error
	// 	toggleResourceCardForDomesticTrade(userID primitive.ObjectID, resourceCardID primitive.ObjectID) error
	// 	offerDomesticTrade(userID primitive.ObjectID, playerID primitive.ObjectID) error
	// 	confirmDomesticTrade(userID primitive.ObjectID) error
	// 	playKnightCard(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID *primitive.ObjectID) error
	// 	playRoadBuildingCard(userID primitive.ObjectID, roadIDs []primitive.ObjectID) error
	// 	playYearOfPlentyCard(userID primitive.ObjectID, resourceCardTypes []ResourceCardType) error
	// 	playMonopolyCard(userID primitive.ObjectID, resourceCardType ResourceCardType) error

}
