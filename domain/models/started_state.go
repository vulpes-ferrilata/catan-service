package models

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type startedState struct {
	game *Game
}

func (s startedState) getPhase() phase {
	switch s.game.phase {
	case Setup:
		return &setupPhase{s.game}
	case ResourceProduction:
		return &resourceProductionPhase{s.game}
	case ResourceDiscard:
		return &resourceDiscardPhase{s.game}
	case Robbing:
		return &robbingPhase{s.game}
	case ResourceConsumption:
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
	if err := s.getPhase().buildSettlementAndRoad(userID, landID, pathID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) rollDices(userID primitive.ObjectID) error {
	if err := s.getPhase().rollDices(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) discardResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	if err := s.getPhase().discardResourceCards(userID, resourceCardIDs); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	if err := s.getPhase().moveRobber(userID, terrainID, playerID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) endTurn(userID primitive.ObjectID) error {
	if err := s.getPhase().endTurn(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	if err := s.getPhase().buildSettlement(userID, landID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	if err := s.getPhase().buildRoad(userID, pathID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	if err := s.getPhase().upgradeCity(userID, constructionID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) buyDevelopmentCard(userID primitive.ObjectID) error {
	if err := s.getPhase().buyDevelopmentCard(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) toggleResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	if err := s.getPhase().toggleResourceCards(userID, resourceCardIDs); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) maritimeTrade(userID primitive.ObjectID, demandingResourceCardType ResourceCardType) error {
	if err := s.getPhase().maritimeTrade(userID, demandingResourceCardType); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) sendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	if err := s.getPhase().sendTradeOffer(userID, playerID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) confirmTradeOffer(userID primitive.ObjectID) error {
	if err := s.getPhase().confirmTradeOffer(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) cancelTradeOffer(userID primitive.ObjectID) error {
	if err := s.getPhase().cancelTradeOffer(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s startedState) playKnightCard(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	if err := s.game.useDevelopmentCard(Knight); err != nil {
		return errors.WithStack(err)
	}

	terrain, isExists := slices.Find(func(terrain *Terrain) bool {
		return terrain.id == terrainID
	}, s.game.terrains)
	if !isExists {
		return errors.WithStack(app_errors.ErrTerrainNotFound)
	}

	var player *Player
	if playerID != primitive.NilObjectID {
		player, isExists = slices.Find(func(player *Player) bool {
			return player.id == playerID
		}, s.game.players)
		if !isExists {
			return errors.WithStack(app_errors.ErrPlayerNotFound)
		}
	}

	if err := s.game.moveRobber(terrain); err != nil {
		return errors.WithStack(err)
	}

	if err := s.game.robPlayer(player); err != nil {
		return errors.WithStack(err)
	}

	if err := s.game.dispatchLargestArmyDevelopment(); err != nil {
		return errors.WithStack(err)
	}

	s.game.calculateScore()

	return nil
}

func (s startedState) playRoadBuildingCard(userID primitive.ObjectID, pathIDs []primitive.ObjectID) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	paths, err := slices.Map(func(pathID primitive.ObjectID) (*Path, error) {
		path, isExists := slices.Find(func(path *Path) bool {
			return path.id == pathID
		}, s.game.paths)
		if !isExists {
			return nil, errors.WithStack(app_errors.ErrPathNotFound)
		}

		return path, nil
	}, pathIDs)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := s.game.useDevelopmentCard(RoadBuilding); err != nil {
		return errors.WithStack(err)
	}

	for _, path := range paths {
		if err := s.game.buildRoad(path); err != nil {
			return errors.WithStack(err)
		}
	}

	if err := s.game.dispatchLongestRoadAchievement(); err != nil {
		return errors.WithStack(err)
	}

	s.game.calculateScore()

	return nil
}

func (s startedState) playYearOfPlentyCard(userID primitive.ObjectID, resourceCardTypes []ResourceCardType) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	if err := s.game.useDevelopmentCard(YearOfPlenty); err != nil {
		return errors.WithStack(err)
	}

	for _, resourceCardType := range resourceCardTypes {
		resourceCard, isExists := slices.Find(func(resourceCard *ResourceCard) bool {
			return resourceCard.resourceCardType == resourceCardType
		}, s.game.resourceCards)
		if !isExists {
			return errors.WithStack(app_errors.ErrGameHasInsufficientResourceCards)
		}

		s.game.resourceCards = slices.Remove(s.game.resourceCards, resourceCard)
		s.game.activePlayer.resourceCards = append(s.game.activePlayer.resourceCards, resourceCard)
	}

	return nil
}

func (s startedState) playMonopolyCard(userID primitive.ObjectID, resourceCardType ResourceCardType) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	if err := s.game.useDevelopmentCard(Monopoly); err != nil {
		return errors.WithStack(err)
	}

	resourceCards := make([]*ResourceCard, 0)

	for _, player := range s.game.players {
		for _, resourceCard := range player.resourceCards {
			if resourceCard.resourceCardType == resourceCardType {
				player.resourceCards = slices.Remove(player.resourceCards, resourceCard)
				resourceCards = append(resourceCards, resourceCard)
			}
		}
	}

	s.game.activePlayer.resourceCards = append(s.game.activePlayer.resourceCards, resourceCards...)

	return nil
}
