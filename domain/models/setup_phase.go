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
	activePlayer, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, s.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if !activePlayer.isActive {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	land, isExists := slices.Find(func(land *Land) bool {
		return land.id == landID
	}, s.game.lands)
	if !isExists {
		return errors.WithStack(app_errors.ErrSelectedLandHasBeenOccupied)
	}

	path, isExists := slices.Find(func(path *Path) bool {
		return path.id == pathID
	}, s.game.paths)
	if !isExists {
		return errors.WithStack(app_errors.ErrSelectedPathHasBeenOccupied)
	}

	for _, player := range s.game.players {
		for _, construction := range player.constructions {
			if construction.land != nil && construction.land.hexCorner.IsAdjacentWithHexCorner(land.hexCorner) {
				return errors.WithStack(app_errors.ErrNearbyLandsMustBeVacant)
			}
		}
	}

	if !path.hexEdge.IsAdjacentWithHexCorner(land.hexCorner) {
		return errors.WithStack(app_errors.ErrSelectedLandAndPathMustBeAdjacent)
	}

	settlement, isExists := slices.Find(func(construction *Construction) bool {
		return construction.land == nil && construction.constructionType == Settlement
	}, activePlayer.constructions)
	if !isExists {
		return errors.WithStack(app_errors.ErrYouHaveRunOutOfSettlements)
	}

	road, isExists := slices.Find(func(road *Road) bool {
		return road.path == nil
	}, activePlayer.roads)
	if !isExists {
		return errors.WithStack(app_errors.ErrYouHaveRunOutOfRoads)
	}

	//build settlement and road
	s.game.lands = slices.Remove(s.game.lands, land)
	s.game.paths = slices.Remove(s.game.paths, path)
	settlement.land = land
	road.path = path

	//dispatch resources
	if s.game.turn == 2 {
		adjacentHexes := FindAdjacentHexesFromHexCorner(land.hexCorner)

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
				activePlayer.resourceCards = append(activePlayer.resourceCards, resourceCard)
			}
		}
	}

	for _, player := range s.game.players {
		player.isActive = false
	}

	switch s.game.turn {
	case 1:
		nextPlayer, isExists := slices.Find(func(player *Player) bool {
			return player.turnOrder == activePlayer.turnOrder+1
		}, s.game.players)
		if !isExists {
			s.game.turn++
			activePlayer.isActive = true
			return nil
		}
		nextPlayer.isActive = true
	case 2:
		nextPlayer, isExists := slices.Find(func(player *Player) bool {
			return player.turnOrder == activePlayer.turnOrder-1
		}, s.game.players)
		if !isExists {
			s.game.turn++
			activePlayer.isActive = true
			return nil
		}
		nextPlayer.isActive = true
	}

	return nil
}

func (s setupPhase) rollDices(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (s setupPhase) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (s setupPhase) endTurn(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}
