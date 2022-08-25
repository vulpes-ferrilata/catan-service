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
