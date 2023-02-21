package models

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/app_errors"
	"github.com/vulpes-ferrilata/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type setupPhase struct {
	game *Game
}

func (s setupPhase) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	if s.game.activePlayer.userID != userID {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	land, err := slices.Find(func(land *Land) (bool, error) {
		return land.id == landID, nil
	}, s.game.lands...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrLandNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	path, err := slices.Find(func(path *Path) (bool, error) {
		return path.id == pathID, nil
	}, s.game.paths...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrPathNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	isLandAdjacentToAnyConstruction, err := slices.Any(func(player *Player) (bool, error) {
		return slices.Any(func(construction *Construction) (bool, error) {
			isAdjacentWithHexCorner, err := construction.land.hexCorner.isAdjacentWithHexCorner(land.hexCorner)
			if err != nil {
				return false, errors.WithStack(err)
			}
			return construction.land != nil && isAdjacentWithHexCorner, nil
		}, player.constructions...)
	}, s.game.getAllPlayers()...)
	if err != nil {
		return errors.WithStack(err)
	}
	if isLandAdjacentToAnyConstruction {
		return errors.WithStack(app_errors.ErrNearbyLandsMustBeVacant)
	}

	isAdjacentWithHexCorner, err := path.hexEdge.isAdjacentWithHexCorner(land.hexCorner)
	if err != nil {
		return errors.WithStack(err)
	}
	if !isAdjacentWithHexCorner {
		return errors.WithStack(app_errors.ErrSelectedLandAndPathMustBeAdjacent)
	}

	settlement, err := slices.Find(func(construction *Construction) (bool, error) {
		return construction.land == nil && construction.constructionType == Settlement, nil
	}, s.game.activePlayer.constructions...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrYouRunOutOfSettlements)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	road, err := slices.Find(func(road *Road) (bool, error) {
		return road.path == nil, nil
	}, s.game.activePlayer.roads...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrYouRunOutOfRoads)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	//build settlement and road
	s.game.lands = slices.Remove(s.game.lands, land)
	s.game.paths = slices.Remove(s.game.paths, path)
	settlement.land = land
	road.path = path

	//dispatch resources
	if s.game.turn == 2 {
		adjacentHexes := findAdjacentHexesFromHexCorner(land.hexCorner)

		terrains, err := slices.Filter(func(terrain *Terrain) (bool, error) {
			_, err := slices.Find(func(adjacentHex Hex) (bool, error) {
				return adjacentHex == terrain.hex, nil
			}, adjacentHexes...)
			if errors.Is(err, slices.ErrNoMatchFound) {
				return false, nil
			}
			if err != nil {
				return false, errors.WithStack(err)
			}

			return terrain.terrainType != Desert, nil
		}, s.game.terrains...)
		if err != nil {
			return errors.WithStack(err)
		}

		for _, terrain := range terrains {
			resourceCard, err := slices.Find(func(resourceCard *ResourceCard) (bool, error) {
				switch terrain.terrainType {
				case Forest:
					return resourceCard.resourceCardType == Lumber, nil
				case Hill:
					return resourceCard.resourceCardType == Brick, nil
				case Pasture:
					return resourceCard.resourceCardType == Wool, nil
				case Field:
					return resourceCard.resourceCardType == Grain, nil
				case Mountain:
					return resourceCard.resourceCardType == Ore, nil
				default:
					return false, nil
				}
			}, s.game.resourceCards...)
			if errors.Is(err, slices.ErrNoMatchFound) {
				continue
			}
			if err != nil {
				return errors.WithStack(err)
			}

			s.game.resourceCards = slices.Remove(s.game.resourceCards, resourceCard)
			s.game.activePlayer.resourceCards = append(s.game.activePlayer.resourceCards, resourceCard)
		}
	}

	switch s.game.turn {
	case 1:
		nextPlayer, err := slices.Find(func(p *Player) (bool, error) {
			return p.turnOrder == s.game.activePlayer.turnOrder+1, nil
		}, s.game.players...)
		if errors.Is(err, slices.ErrNoMatchFound) {
			s.game.turn++
			return nil
		}
		if err != nil {
			return errors.WithStack(err)
		}

		*s.game.activePlayer, *nextPlayer = *nextPlayer, *s.game.activePlayer //swap pointer
	case 2:
		nextPlayer, err := slices.Find(func(p *Player) (bool, error) {
			return p.turnOrder == s.game.activePlayer.turnOrder-1, nil
		}, s.game.players...)
		if errors.Is(err, slices.ErrNoMatchFound) {
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
