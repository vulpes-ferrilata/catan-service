package models

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type startedState struct {
	game *Game
}

func (s startedState) getPhase() phase {
	if s.game.turn == 1 || s.game.turn == 2 {
		return &setupPhase{s.game}
	}

	if !s.game.isRolledDices {
		return &resourceProductionPhase{s.game}
	}

	if s.game.isRolledDices {
		return &resourceConsumptionPhase{s.game}
	}

	return nil
}

func (s startedState) newPlayer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyStarted)
}

func (s startedState) startGame(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameAlreadyStarted)
}

func (s startedState) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	phase := s.getPhase()

	if err := phase.buildSettlementAndRoad(userID, landID, pathID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) rollDices(userID primitive.ObjectID) error {
	phase := s.getPhase()

	if err := phase.rollDices(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	currentPhase := s.getPhase()

	if err := currentPhase.moveRobber(userID, terrainID, playerID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) endTurn(userID primitive.ObjectID) error {
	currentPhase := s.getPhase()

	if err := currentPhase.endTurn(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
