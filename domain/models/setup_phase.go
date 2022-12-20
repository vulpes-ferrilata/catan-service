package models

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type setupPhase struct {
	game *Game
}

func (s setupPhase) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	land, isExists := slices.Find(func(land *Land) bool {
		return land.id == landID
	}, s.game.lands)
	if !isExists {
		return errors.WithStack(app_errors.ErrLandNotFound)
	}

	path, isExists := slices.Find(func(path *Path) bool {
		return path.id == pathID
	}, s.game.paths)
	if !isExists {
		return errors.WithStack(app_errors.ErrPathNotFound)
	}

	isLandAdjacentToAnyConstruction := slices.Any(func(player *Player) bool {
		return slices.Any(func(construction *Construction) bool {
			return construction.land != nil && construction.land.hexCorner.isAdjacentWithHexCorner(land.hexCorner)
		}, player.constructions)
	}, s.game.getAllPlayers())
	if isLandAdjacentToAnyConstruction {
		return errors.WithStack(app_errors.ErrNearbyLandsMustBeVacant)
	}

	if !path.hexEdge.isAdjacentWithHexCorner(land.hexCorner) {
		return errors.WithStack(app_errors.ErrSelectedLandAndPathMustBeAdjacent)
	}

	settlement, isExists := slices.Find(func(construction *Construction) bool {
		return construction.land == nil && construction.constructionType == Settlement
	}, s.game.activePlayer.constructions)
	if !isExists {
		return errors.WithStack(app_errors.ErrYouRunOutOfSettlements)
	}

	road, isExists := slices.Find(func(road *Road) bool {
		return road.path == nil
	}, s.game.activePlayer.roads)
	if !isExists {
		return errors.WithStack(app_errors.ErrYouRunOutOfRoads)
	}

	//build settlement and road
	s.game.lands = slices.Remove(s.game.lands, land)
	s.game.paths = slices.Remove(s.game.paths, path)
	settlement.land = land
	road.path = path

	//dispatch resources
	if s.game.turn == 2 {
		adjacentHexes := findAdjacentHexesFromHexCorner(land.hexCorner)

		terrains := slices.Filter(func(terrain *Terrain) bool {
			return slices.Contains(adjacentHexes, terrain.hex) && terrain.terrainType != Desert
		}, s.game.terrains)

		for _, terrain := range terrains {
			resourceCard, isExists := slices.Find(func(resourceCard *ResourceCard) bool {
				switch terrain.terrainType {
				case Forest:
					return resourceCard.resourceCardType == Lumber
				case Hill:
					return resourceCard.resourceCardType == Brick
				case Pasture:
					return resourceCard.resourceCardType == Wool
				case Field:
					return resourceCard.resourceCardType == Grain
				case Mountain:
					return resourceCard.resourceCardType == Ore
				}

				return false
			}, s.game.resourceCards)
			if isExists {
				s.game.resourceCards = slices.Remove(s.game.resourceCards, resourceCard)
				s.game.activePlayer.resourceCards = append(s.game.activePlayer.resourceCards, resourceCard)
			}
		}
	}

	s.game.calculateScore()

	switch s.game.turn {
	case 1:
		nextPlayer, isExists := slices.Find(func(p *Player) bool {
			return p.turnOrder == s.game.activePlayer.turnOrder+1
		}, s.game.players)
		if !isExists {
			s.game.turn++
			return nil
		}

		*s.game.activePlayer, *nextPlayer = *nextPlayer, *s.game.activePlayer //swap pointer
	case 2:
		nextPlayer, isExists := slices.Find(func(p *Player) bool {
			return p.turnOrder == s.game.activePlayer.turnOrder-1
		}, s.game.players)
		if !isExists {
			s.game.turn++

			s.game.phase = ResourceProduction

			return nil
		}

		*s.game.activePlayer, *nextPlayer = *nextPlayer, *s.game.activePlayer //swap pointer
	}

	return nil
}

func (s setupPhase) rollDices(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) discardResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) endTurn(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) buyDevelopmentCard(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) toggleResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) maritimeTrade(userID primitive.ObjectID, resourceCardType ResourceCardType, demandingResourceCardType ResourceCardType) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) sendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) confirmTradeOffer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}

func (s setupPhase) cancelTradeOffer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInSetupPhase)
}
