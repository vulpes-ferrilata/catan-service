package models

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type finishedState struct {
	game *Game
}

func (f finishedState) newPlayer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) startGame(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) rollDices(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) endTurn(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) buyDevelopmentCard(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) toggleResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) maritimeTrade(userID primitive.ObjectID, demandingResourceCardType ResourceCardType) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) sendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) confirmTradeOffer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) cancelTradeOffer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) playKnightCard(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) playRoadBuildingCard(userID primitive.ObjectID, pathIDs []primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) playYearOfPlentyCard(userID primitive.ObjectID, resourceCardTypes []ResourceCardType) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}

func (f finishedState) playMonopolyCard(userID primitive.ObjectID, resourceCardType ResourceCardType) error {
	return errors.WithStack(app_errors.ErrGameAlreadyFinished)
}
