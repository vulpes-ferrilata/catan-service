package models

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type robbingPhase struct {
	game *Game
}

func (r robbingPhase) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) rollDices(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	if r.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	terrain, isExists := slices.Find(func(terrain *Terrain) bool {
		return terrain.id == terrainID
	}, r.game.terrains)
	if !isExists {
		return errors.WithStack(app_errors.ErrTerrainNotFound)
	}

	var player *Player
	if playerID != primitive.NilObjectID {
		player, isExists = slices.Find(func(player *Player) bool {
			return player.id == playerID
		}, r.game.players)
		if !isExists {
			return errors.WithStack(app_errors.ErrPlayerNotFound)
		}
	}

	if err := r.game.moveRobber(terrain); err != nil {
		return errors.WithStack(err)
	}

	if err := r.game.robPlayer(player); err != nil {
		return errors.WithStack(err)
	}

	r.game.phase = ResourceConsumption

	return nil
}

func (r robbingPhase) endTurn(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) buyDevelopmentCard(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) toggleResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) maritimeTrade(userID primitive.ObjectID, demandingResourceCardType ResourceCardType) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) sendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) confirmTradeOffer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}

func (r robbingPhase) cancelTradeOffer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInRobbingPhase)
}
