package models

import (
	"math/rand"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type resourceConsumptionPhase struct {
	game *Game
}

func (r resourceConsumptionPhase) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceConsumptionPhase) rollDices(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInCurrentPhase)
}

func (r resourceConsumptionPhase) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	activePlayer, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if !activePlayer.isActive {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	//TODO: robber need to moving
	if !r.game.robber.isMoving {
		return errors.WithStack(app_errors.ErrRobberIsNotAvailableToMove)
	}

	r.game.robber.isMoving = false

	if r.game.robber.terrainID == terrainID {
		return errors.WithStack(app_errors.ErrRobberMustBeMovedToOtherTerrain)
	}

	r.game.robber.terrainID = terrainID

	terrain, isExists := slices.Find(func(terrain *Terrain) bool {
		return terrain.id == terrainID
	}, r.game.terrains)
	if !isExists {
		return errors.WithStack(app_errors.ErrTerrainNotFound)
	}

	if playerID == primitive.NilObjectID {
		for _, player := range r.game.players {
			if player.userID != userID {
				canBeRob := slices.Some(func(construction *Construction) bool {
					return construction.land != nil && construction.land.hexCorner.IsAdjacentWithHex(terrain.hex)
				}, player.constructions)
				if canBeRob {
					return errors.WithStack(app_errors.ErrYouMustRobPlayerWhoHaveConstructionNextToRobber)
				}
			}
		}
	}

	robbingPlayer, isExists := slices.Find(func(player *Player) bool {
		return player.id == playerID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}

	if activePlayer == robbingPlayer {
		return errors.WithStack(app_errors.ErrYouCannotRobYourself)
	}

	canBeRob := slices.Some(func(construction *Construction) bool {
		return construction.land != nil && construction.land.hexCorner.IsAdjacentWithHex(terrain.hex)
	}, robbingPlayer.constructions)
	if !canBeRob {
		return errors.WithStack(app_errors.ErrSelectedPlayerMustHaveConstructionNextToRobber)
	}

	if len(robbingPlayer.resourceCards) > 0 {
		resourceCardIdx := rand.Intn(len(robbingPlayer.resourceCards))
		resourceCard := robbingPlayer.resourceCards[resourceCardIdx]
		robbingPlayer.resourceCards = slices.Remove(robbingPlayer.resourceCards, resourceCard)
		activePlayer.resourceCards = append(activePlayer.resourceCards, resourceCard)
	}

	return nil
}

func (r resourceConsumptionPhase) endTurn(userID primitive.ObjectID) error {
	activePlayer, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, r.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if !activePlayer.isActive {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	if !r.game.isRolledDices {
		return errors.WithStack(app_errors.ErrYouMustRollDices)
	}

	r.game.isRolledDices = false

	for _, player := range r.game.players {
		for _, resourceCard := range player.resourceCards {
			resourceCard.isSelected = false
		}

		for _, developmentCard := range player.developmentCards {
			if developmentCard.status == Disable {
				developmentCard.status = Enable
			}
		}

		player.isActive = false
	}

	nextPlayer, isExists := slices.Find(func(player *Player) bool {
		return player.turnOrder == activePlayer.turnOrder+1
	}, r.game.players)
	if !isExists {
		r.game.turn++
		nextPlayer, isExists := slices.Find(func(player *Player) bool {
			return player.turnOrder == 1
		}, r.game.players)
		if !isExists {
			return errors.WithStack(app_errors.ErrPlayerNotFound)
		}
		nextPlayer.isActive = true
		return nil
	}
	nextPlayer.isActive = true

	return nil
}

// func (r resourceConsumptionPhase) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
// 	currentPlayer, err := r.game.GetCurrentPlayer()
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}
// 	if currentPlayer.userID != userID {
// 		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
// 	}

// 	lumberResourceCard, isExists := currentPlayer.resourceCards.Find(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Brick
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	currentPlayer.resourceCards.Remove(lumberResourceCard)
// 	r.game.resourceCards.Add(lumberResourceCard)

// 	brickResourceCard, isExists := currentPlayer.resourceCards.Find(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Brick
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	currentPlayer.resourceCards.Remove(brickResourceCard)
// 	r.game.resourceCards.Add(brickResourceCard)

// 	woolResourceCard, isExists := currentPlayer.resourceCards.Find(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Wool
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	currentPlayer.resourceCards.Remove(woolResourceCard)
// 	r.game.resourceCards.Add(woolResourceCard)

// 	grainResourceCard, isExists := currentPlayer.resourceCards.Find(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Grain
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	currentPlayer.resourceCards.Remove(grainResourceCard)
// 	r.game.resourceCards.Add(grainResourceCard)

// 	land, isExists := r.game.lands.Find(func(land *Land) bool {
// 		return land.id == landID
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrSelectedLandHasBeenOccupied)
// 	}

// 	if err := r.game.players.ForEach(func(player *Player) error {
// 		return player.constructions.ForEach(func(construction *Construction) error {
// 			if construction.land != nil && construction.land.hexCorner.IsAdjacentWithHexCorner(land.hexCorner) {
// 				return errors.WithStack(app_errors.ErrNearbyLandsMustBeVacant)
// 			}

// 			return nil
// 		})
// 	}); err != nil {
// 		return errors.WithStack(err)
// 	}

// 	if isAdjacentWithRoad := currentPlayer.roads.Some(func(road *Road) bool {
// 		if road.path != nil && road.path.hexEdge.IsAdjacentWithHexCorner(land.hexCorner) {
// 			return true
// 		}

// 		return false
// 	}); !isAdjacentWithRoad {
// 		return errors.WithStack(app_errors.ErrSelectedLandMustBeAdjacentWithYourRoad)
// 	}

// 	settlement, isExists := currentPlayer.constructions.Find(func(construction *Construction) bool {
// 		return construction.land == nil && construction.constructionType == Settlement
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveRunOutOfSettlements)
// 	}

// 	r.game.lands.Remove(land)
// 	settlement.land = land

// 	//TODO: calculate score

// 	return nil
// }

// func (r resourceConsumptionPhase) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
// 	currentPlayer, err := r.game.GetCurrentPlayer()
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}
// 	if currentPlayer.userID != userID {
// 		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
// 	}

// 	lumberResourceCard, isExists := currentPlayer.resourceCards.Find(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Brick
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	currentPlayer.resourceCards.Remove(lumberResourceCard)
// 	r.game.resourceCards.Add(lumberResourceCard)

// 	brickResourceCard, isExists := currentPlayer.resourceCards.Find(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Brick
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	currentPlayer.resourceCards.Remove(brickResourceCard)
// 	r.game.resourceCards.Add(brickResourceCard)

// 	path, isExists := r.game.paths.Find(func(path *Path) bool {
// 		return path.id == pathID
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrSelectedPathHasBeenOccupied)
// 	}

// 	isAdjacentWithConstruction := currentPlayer.constructions.Some(func(construction *Construction) bool {
// 		if construction.land != nil && construction.land.hexCorner.IsAdjacentWithHexEdge(path.hexEdge) {
// 			return true
// 		}

// 		return false
// 	})
// 	isAdjacentWithRoad := currentPlayer.roads.Some(func(road *Road) bool {
// 		if road.path != nil && road.path.hexEdge.IsAdjacentWithHexEdge(path.hexEdge) {
// 			return true
// 		}

// 		return false
// 	})
// 	if !isAdjacentWithConstruction && !isAdjacentWithRoad {
// 		return errors.WithStack(app_errors.ErrSelectedPathMustBeAdjacentWithYourConstructionOrRoad)
// 	}

// 	road, isExists := currentPlayer.roads.Find(func(road *Road) bool {
// 		return road.path == nil
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveRunOutOfRoads)
// 	}

// 	r.game.paths.Remove(path)
// 	road.path = path

// 	var longestRoadAchievement *Achievement

// 	r.game.achievements.ForEach(func(achievement *Achievement) error {
// 		if achievement.achievementType == LongestRoad {
// 			longestRoadAchievement = achievement
// 		}

// 		return nil
// 	})

// 	if err := r.game.players.ForEach(func(player *Player) error {
// 		return player.achievements.ForEach(func(achievement *Achievement) error {
// 			if achievement.achievementType == LongestRoad {
// 				longestRoadAchievement = achievement
// 			}

// 			return nil
// 		})
// 	}); err != nil {
// 		return errors.WithStack(err)
// 	}

// 	if longestRoadAchievement == nil {
// 		return errors.New("longest road achievement is not exists")
// 	}

// 	r.game.achievements.Remove(longestRoadAchievement)
// 	if err := r.game.players.ForEach(func(player *Player) error {
// 		player.achievements.Remove(longestRoadAchievement)

// 		return nil
// 	}); err != nil {
// 		return errors.WithStack(err)
// 	}

// 	longestRoad := 0
// 	var longestRoadPlayer *Player
// 	r.game.players.ForEach(func(player *Player) error {
// 		result := r.calculateLongestRoad(player.roads.ToSlice())
// 		if result == longestRoad {
// 			longestRoadPlayer = nil
// 		}
// 		if result > longestRoad {
// 			longestRoad = result
// 			longestRoadPlayer = player
// 		}

// 		return nil
// 	})
// 	if longestRoad >= 5 && longestRoadPlayer == nil {
// 		r.game.achievements.Add(longestRoadAchievement)
// 	} else {
// 		longestRoadPlayer.achievements.Add(longestRoadAchievement)
// 	}

// 	//TODO: calculate score

// 	return nil
// }

// func (r resourceConsumptionPhase) calculateLongestRoad(roads []*Road) int {
// 	longestRoad := 0

// 	for idx, road := range roads {
// 		remainRoads := append(roads[:idx], roads[idx+1:]...)
// 		result := r.calculateLongestRoadFromCurrentRoad(road, remainRoads)
// 		if result > longestRoad {
// 			longestRoad = result
// 		}
// 	}

// 	return longestRoad
// }

// func (r resourceConsumptionPhase) calculateLongestRoadFromCurrentRoad(currentRoad *Road, otherRoads []*Road) int {
// 	longestRoad := 0

// 	for idx, otherRoad := range otherRoads {
// 		if currentRoad.path != nil && otherRoad.path != nil && currentRoad.path.hexEdge.IsAdjacentWithHexEdge(otherRoad.path.hexEdge) {
// 			remainRoads := append(otherRoads[:idx], otherRoads[idx+1:]...)
// 			result := r.calculateLongestRoadFromCurrentRoad(otherRoad, remainRoads)
// 			if result > longestRoad {
// 				longestRoad = result
// 			}
// 		}
// 	}

// 	return 1 + longestRoad
// }

// func (r resourceConsumptionPhase) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
// 	currentPlayer, err := r.game.GetCurrentPlayer()
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}
// 	if currentPlayer.userID != userID {
// 		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
// 	}

// 	grainResourceCards := currentPlayer.resourceCards.Filter(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Grain
// 	}).ToSlice()
// 	if len(grainResourceCards) < 2 {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	grainResourceCards = grainResourceCards[:2]
// 	currentPlayer.resourceCards.Remove(grainResourceCards...)
// 	r.game.resourceCards.Add(grainResourceCards...)

// 	oreResourceCards := currentPlayer.resourceCards.Filter(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Ore
// 	}).ToSlice()
// 	if len(grainResourceCards) < 3 {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	oreResourceCards = oreResourceCards[:3]
// 	currentPlayer.resourceCards.Remove(oreResourceCards...)
// 	r.game.resourceCards.Add(oreResourceCards...)

// 	construction, isExists := currentPlayer.constructions.Find(func(construction *Construction) bool {
// 		return construction.id == constructionID
// 	})
// 	if !isExists {
// 		return errors.New("construction is not exists")
// 	}
// 	if construction.constructionType == City {
// 		return errors.WithStack(app_errors.ErrSelectedConstructionAlreadyUpgraded)
// 	}

// 	land := construction.land

// 	if land == nil {
// 		return errors.WithStack(app_errors.SelectedConstructionDoesNotBelongToAnyLand)
// 	}

// 	city, isExists := currentPlayer.constructions.Find(func(construction *Construction) bool {
// 		return construction.constructionType == City && construction.land == nil
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveRunOutOfCities)
// 	}

// 	construction.land = nil
// 	city.land = land

// 	//TODO: calculate score

// 	return nil
// }

// func (r resourceConsumptionPhase) buyDevelopmentCard(userID primitive.ObjectID) error {
// 	currentPlayer, err := r.game.GetCurrentPlayer()
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}
// 	if currentPlayer.userID != userID {
// 		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
// 	}

// 	woolResourceCard, isExists := currentPlayer.resourceCards.Find(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Wool
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	currentPlayer.resourceCards.Remove(woolResourceCard)
// 	r.game.resourceCards.Add(woolResourceCard)

// 	grainResourceCard, isExists := currentPlayer.resourceCards.Find(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Grain
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	currentPlayer.resourceCards.Remove(grainResourceCard)
// 	r.game.resourceCards.Add(grainResourceCard)

// 	oreResourceCard, isExists := currentPlayer.resourceCards.Find(func(resourceCard *ResourceCard) bool {
// 		return resourceCard.resourceCardType == Ore
// 	})
// 	if !isExists {
// 		return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
// 	}
// 	currentPlayer.resourceCards.Remove(oreResourceCard)
// 	r.game.resourceCards.Add(oreResourceCard)

// 	if r.game.developmentCards.Length() == 0 {
// 		return errors.WithStack(app_errors.ErrGameHasRunOutOfDevelopmentCards)
// 	}

// 	developmentCards := r.game.developmentCards.ToSlice()
// 	rand.Shuffle(len(developmentCards), func(i, j int) { developmentCards[i], developmentCards[j] = developmentCards[j], developmentCards[i] })
// 	developmentCards = developmentCards[:1]
// 	r.game.developmentCards.Remove(developmentCards...)
// 	currentPlayer.developmentCards.Add(developmentCards...)

// 	//TODO: calculate score

// 	return nil
// }
