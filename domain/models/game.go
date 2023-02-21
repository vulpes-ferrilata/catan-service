package models

import (
	"math/rand"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/app_errors"
	"github.com/vulpes-ferrilata/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	aggregateRoot
	status           GameStatus
	phase            GamePhase
	turn             int
	activePlayer     *Player
	players          []*Player
	dices            []*Dice
	achievements     []*Achievement
	resourceCards    []*ResourceCard
	developmentCards []*DevelopmentCard
	terrains         []*Terrain
	lands            []*Land
	paths            []*Path
}

func (g Game) GetStatus() GameStatus {
	return g.status
}

func (g Game) GetPhase() GamePhase {
	return g.phase
}

func (g Game) GetTurn() int {
	return g.turn
}

func (g Game) GetActivePlayer() *Player {
	return g.activePlayer
}

func (g Game) GetPlayers() []*Player {
	return g.players
}

func (g Game) GetDices() []*Dice {
	return g.dices
}

func (g Game) GetAchievements() []*Achievement {
	return g.achievements
}

func (g Game) GetResourceCards() []*ResourceCard {
	return g.resourceCards
}

func (g Game) GetDevelopmentCards() []*DevelopmentCard {
	return g.developmentCards
}

func (g Game) GetTerrains() []*Terrain {
	return g.terrains
}

func (g Game) GetLands() []*Land {
	return g.lands
}

func (g Game) GetPaths() []*Path {
	return g.paths
}

func (g *Game) getState() state {
	switch g.status {
	case Waiting:
		return &waitingState{g}
	case Started:
		return &startedState{g}
	case Finished:
		return &finishedState{g}
	}

	return nil
}

func (g Game) getAllPlayers() []*Player {
	allPlayers := make([]*Player, 0)
	if g.activePlayer != nil {
		allPlayers = append(allPlayers, g.activePlayer)
	}
	allPlayers = append(allPlayers, g.players...)

	return allPlayers
}

func (g *Game) NewPlayer(userID primitive.ObjectID) error {
	if err := g.getState().newPlayer(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) Start(userID primitive.ObjectID) error {
	if err := g.getState().startGame(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) BuildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	if err := g.getState().buildSettlementAndRoad(userID, landID, pathID); err != nil {
		return errors.WithStack(err)
	}

	g.calculateScore()

	return nil
}

func (g *Game) RollDices(userID primitive.ObjectID) error {
	if err := g.getState().rollDices(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) DiscardResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	if err := g.getState().discardResourceCards(userID, resourceCardIDs); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) MoveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	if err := g.getState().moveRobber(userID, terrainID, playerID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) EndTurn(userID primitive.ObjectID) error {
	if err := g.getState().endTurn(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) BuildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	if err := g.getState().buildSettlement(userID, landID); err != nil {
		return errors.WithStack(err)
	}

	if err := g.dispatchLongestRoadAchievement(); err != nil {
		return errors.WithStack(err)
	}

	g.calculateScore()

	return nil
}

func (g *Game) BuildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	if err := g.getState().buildRoad(userID, pathID); err != nil {
		return errors.WithStack(err)
	}

	if err := g.dispatchLongestRoadAchievement(); err != nil {
		return errors.WithStack(err)
	}

	g.calculateScore()

	return nil
}

func (g *Game) UpgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	if err := g.getState().upgradeCity(userID, constructionID); err != nil {
		return errors.WithStack(err)
	}

	g.calculateScore()

	return nil
}

func (g *Game) BuyDevelopmentCard(userID primitive.ObjectID) error {
	if err := g.getState().buyDevelopmentCard(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) ToggleResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	if err := g.getState().toggleResourceCards(userID, resourceCardIDs); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) MaritimeTrade(userID primitive.ObjectID, resourceCardType ResourceCardType, demandingResourceCardType ResourceCardType) error {
	if err := g.getState().maritimeTrade(userID, resourceCardType, demandingResourceCardType); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) SendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	if err := g.getState().sendTradeOffer(userID, playerID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) ConfirmTradeOffer(userID primitive.ObjectID) error {
	if err := g.getState().confirmTradeOffer(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) CancelTradeOffer(userID primitive.ObjectID) error {
	if err := g.getState().cancelTradeOffer(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) PlayKnightCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	if err := g.getState().playKnightCard(userID, developmentCardID, terrainID, playerID); err != nil {
		return errors.WithStack(err)
	}

	if err := g.dispatchLargestArmyDevelopment(); err != nil {
		return errors.WithStack(err)
	}

	g.calculateScore()

	return nil
}

func (g *Game) PlayRoadBuildingCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, pathIDs []primitive.ObjectID) error {
	if err := g.getState().playRoadBuildingCard(userID, developmentCardID, pathIDs); err != nil {
		return errors.WithStack(err)
	}

	if err := g.dispatchLongestRoadAchievement(); err != nil {
		return errors.WithStack(err)
	}

	g.calculateScore()

	return nil
}

func (g *Game) PlayYearOfPlentyCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, demandingResourceCardTypes []ResourceCardType) error {
	if err := g.getState().playYearOfPlentyCard(userID, developmentCardID, demandingResourceCardTypes); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) PlayMonopolyCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID, demandingResourceCardType ResourceCardType) error {
	if err := g.getState().playMonopolyCard(userID, developmentCardID, demandingResourceCardType); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) PlayVictoryPointCard(userID primitive.ObjectID, developmentCardID primitive.ObjectID) error {
	if err := g.getState().playVictoryPointCard(userID, developmentCardID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) useResourceCards(resourceCardTypes ...ResourceCardType) error {
	for _, resourceCardType := range resourceCardTypes {
		resourceCard, err := slices.Find(func(resourceCard *ResourceCard) (bool, error) {
			return resourceCard.resourceCardType == resourceCardType, nil
		}, g.activePlayer.resourceCards...)
		if errors.Is(err, slices.ErrNoMatchFound) {
			return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
		}
		if err != nil {
			return errors.WithStack(err)
		}

		g.activePlayer.resourceCards = slices.Remove(g.activePlayer.resourceCards, resourceCard)
		g.resourceCards = append(g.resourceCards, resourceCard)
	}

	return nil
}

func (g *Game) buildSettlement(land *Land) error {
	isLandAdjacentToAnyConstruction, err := slices.Any(func(player *Player) (bool, error) {
		return slices.Any(func(construction *Construction) (bool, error) {
			isAdjacentWithHexCorner, err := construction.land.hexCorner.isAdjacentWithHexCorner(land.hexCorner)
			if err != nil {
				return false, errors.WithStack(err)
			}
			return construction.land != nil && isAdjacentWithHexCorner, nil
		}, player.constructions...)
	}, g.players...)
	if err != nil {
		return errors.WithStack(err)
	}
	if isLandAdjacentToAnyConstruction {
		return errors.WithStack(app_errors.ErrNearbyLandsMustBeVacant)
	}

	isLandAdjacentToPlayerRoad, err := slices.Any(func(road *Road) (bool, error) {
		isAdjacentWithHexCorner, err := road.path.hexEdge.isAdjacentWithHexCorner(land.hexCorner)
		if err != nil {
			return false, errors.WithStack(err)
		}

		return road.path != nil && isAdjacentWithHexCorner, nil
	}, g.activePlayer.roads...)
	if err != nil {
		return errors.WithStack(err)
	}
	if !isLandAdjacentToPlayerRoad {
		return errors.WithStack(app_errors.ErrSelectedLandMustBeAdjacentToYourRoad)
	}

	settlement, err := slices.Find(func(construction *Construction) (bool, error) {
		return construction.constructionType == Settlement && construction.land == nil, nil
	}, g.activePlayer.constructions...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrYouRunOutOfSettlements)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	g.lands = slices.Remove(g.lands, land)
	settlement.land = land

	return nil
}

func (g *Game) buildRoad(path *Path) error {
	isAdjacentToConstruction, err := slices.Any(func(construction *Construction) (bool, error) {
		isAdjacentWithHexEdge, err := construction.land.hexCorner.isAdjacentWithHexEdge(path.hexEdge)
		if err != nil {
			return false, errors.WithStack(err)
		}

		return construction.land != nil && isAdjacentWithHexEdge, nil
	}, g.activePlayer.constructions...)
	if err != nil {
		return errors.WithStack(err)
	}

	//if selected path adjacent to your construction, other adjacent land cannot be occupied by other player
	if !isAdjacentToConstruction {
		adjacentHexCorners := findAdjacentHexCornersFromHexEdge(path.hexEdge)

		intersectionHexCorners, err := slices.Filter(func(adjacentHexCorner HexCorner) (bool, error) {
			return slices.Any(func(road *Road) (bool, error) {
				isAdjacentWithHexEdge, err := adjacentHexCorner.isAdjacentWithHexEdge(road.path.hexEdge)
				if err != nil {
					return false, errors.WithStack(err)
				}

				return road.path != nil && isAdjacentWithHexEdge, nil
			}, g.activePlayer.roads...)
		}, adjacentHexCorners...)
		if err != nil {
			return errors.WithStack(err)
		}

		if len(intersectionHexCorners) == 0 {
			return errors.WithStack(app_errors.ErrSelectedPathMustBeAdjacentToYourConstructionOrRoad)
		}

		if len(intersectionHexCorners) == 1 {
			isSelectedPathPassThroughConstructionOfOtherPlayer, err := slices.Any(func(otherPlayer *Player) (bool, error) {
				return slices.Any(func(construction *Construction) (bool, error) {
					return construction.land != nil && construction.land.hexCorner == intersectionHexCorners[0], nil
				}, otherPlayer.constructions...)
			}, g.players...)
			if err != nil {
				return errors.WithStack(err)
			}
			if isSelectedPathPassThroughConstructionOfOtherPlayer {
				return errors.WithStack(app_errors.ErrSelectedPathPassThroughConstructionOfOtherPlayer)
			}
		}

		//if you have roads on both sides of selected path, both sides cannot be occupied
	}

	road, err := slices.Find(func(road *Road) (bool, error) {
		return road.path == nil, nil
	}, g.activePlayer.roads...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrYouRunOutOfRoads)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	g.paths = slices.Remove(g.paths, path)
	road.path = path

	return nil
}

func (g *Game) upgradeConstruction(construction *Construction) error {
	if construction.constructionType == City {
		return errors.WithStack(app_errors.ErrSelectedConstructionAlreadyUpgraded)
	}

	if construction.land == nil {
		return errors.WithStack(app_errors.ErrSelectedConstructionDoesNotBelongToAnyLand)
	}

	city, err := slices.Find(func(construction *Construction) (bool, error) {
		return construction.land == nil && construction.constructionType == City, nil
	}, g.activePlayer.constructions...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrYouRunOutOfCities)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	city.land = construction.land
	construction.land = nil

	return nil
}

func (g *Game) moveRobber(terrain *Terrain) error {
	if terrain.robber != nil {
		return errors.WithStack(app_errors.ErrRobberMustBeMovedToOtherTerrain)
	}

	var robber *Robber

	for _, terrain := range g.terrains {
		if terrain.robber != nil {
			robber = terrain.robber
			terrain.robber = nil
		}
	}

	terrain.robber = robber

	return nil
}

func (g *Game) robPlayer(player *Player) error {
	terrainHasRobber, err := slices.Find(func(terrain *Terrain) (bool, error) {
		return terrain.robber != nil, nil
	}, g.terrains...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrTerrainNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if player == nil {
		hasPlayerCanBeRob, err := slices.Any(func(player *Player) (bool, error) {
			return slices.Any(func(construction *Construction) (bool, error) {
				isAdjacentWithHex, err := construction.land.hexCorner.isAdjacentWithHex(terrainHasRobber.hex)
				if err != nil {
					return false, errors.WithStack(err)
				}

				return construction.land != nil && isAdjacentWithHex, nil
			}, player.constructions...)
		}, g.players...)
		if err != nil {
			return errors.WithStack(err)
		}
		if hasPlayerCanBeRob {
			return errors.WithStack(app_errors.ErrYouMustRobPlayerWhoHasConstructionNextToRobber)
		}

		return nil
	}

	canBeRob, err := slices.Any(func(construction *Construction) (bool, error) {
		isAdjacentWithHex, err := construction.land.hexCorner.isAdjacentWithHex(terrainHasRobber.hex)
		if err != nil {
			return false, errors.WithStack(err)
		}

		return construction.land != nil && isAdjacentWithHex, nil
	}, player.constructions...)
	if err != nil {
		return errors.WithStack(err)
	}
	if !canBeRob {
		return errors.WithStack(app_errors.ErrSelectedPlayerMustHaveConstructionNextToRobber)
	}

	if len(player.resourceCards) > 0 {
		resourceCardIdx := rand.Intn(len(player.resourceCards))
		resourceCard := player.resourceCards[resourceCardIdx]

		player.resourceCards = slices.Remove(player.resourceCards, resourceCard)
		g.activePlayer.resourceCards = append(g.activePlayer.resourceCards, resourceCard)
	}

	return nil
}

func (g *Game) dispatchLongestRoadAchievement() error {
	var longestRoadAchievement *Achievement

	for _, achievement := range g.achievements {
		if achievement.achievementType == LongestRoad {
			longestRoadAchievement = achievement
			g.achievements = slices.Remove(g.achievements, longestRoadAchievement)
		}
	}

	for _, player := range g.getAllPlayers() {
		for _, achievement := range player.achievements {
			if achievement.achievementType == LongestRoad {
				longestRoadAchievement = achievement
				player.achievements = slices.Remove(player.achievements, longestRoadAchievement)
			}
		}
	}

	var playerHasLongestRoad *Player
	maxLongestRoad := 0
	for _, player := range g.getAllPlayers() {
		longestRoad, err := g.calculateLongestRoad(player)
		if err != nil {
			return errors.WithStack(err)
		}

		if longestRoad == maxLongestRoad {
			playerHasLongestRoad = nil
		}

		if longestRoad >= 5 && longestRoad > maxLongestRoad {
			maxLongestRoad = longestRoad
			playerHasLongestRoad = player
		}
	}

	if playerHasLongestRoad != nil {
		playerHasLongestRoad.achievements = append(playerHasLongestRoad.achievements, longestRoadAchievement)
	} else {
		g.achievements = append(g.achievements, longestRoadAchievement)
	}

	return nil
}

func (g Game) calculateLongestRoad(player *Player) (int, error) {
	longestRoad := 0

	for _, road := range player.roads {
		remainRoads := slices.Remove(player.roads, road)
		totalRoads, err := g.calculateLongestRoadFromCurrentRoad(player, road, remainRoads)
		if err != nil {
			return 0, errors.WithStack(err)
		}
		if totalRoads > longestRoad {
			longestRoad = totalRoads
		}
	}

	return longestRoad, nil
}

func (g Game) calculateLongestRoadFromCurrentRoad(player *Player, currentRoad *Road, otherRoads []*Road) (int, error) {
	longestRoad := 0

	if currentRoad.path == nil {
		return 0, nil
	}

	for _, otherRoad := range otherRoads {
		if otherRoad.path == nil {
			continue
		}

		hexCorner, err := findIntersectionHexCornerBetweenTwoHexEdges(currentRoad.path.hexEdge, otherRoad.path.hexEdge)
		if errors.Is(err, slices.ErrNoMatchFound) {
			continue
		}
		if err != nil {
			return 0, errors.WithStack(err)
		}

		otherPlayers := slices.Remove(g.players, player)

		isOtherPlayerHasConstructionBetweenTwoRoads, err := slices.Any(func(otherPlayer *Player) (bool, error) {
			return slices.Any(func(construction *Construction) (bool, error) {
				return construction.land != nil && construction.land.hexCorner == hexCorner, nil
			}, otherPlayer.constructions...)
		}, otherPlayers...)
		if err != nil {
			return 0, errors.WithStack(err)
		}
		if isOtherPlayerHasConstructionBetweenTwoRoads {
			continue
		}

		remainRoads := slices.Remove(otherRoads, otherRoad)
		result, err := g.calculateLongestRoadFromCurrentRoad(player, otherRoad, remainRoads)
		if err != nil {
			return 0, errors.WithStack(err)
		}
		if result > longestRoad {
			longestRoad = result
		}
	}

	return 1 + longestRoad, nil
}

func (g *Game) dispatchLargestArmyDevelopment() error {
	var largestArmyAchievement *Achievement

	for _, achievement := range g.achievements {
		if achievement.achievementType == LargestArmy {
			largestArmyAchievement = achievement
			g.achievements = slices.Remove(g.achievements, largestArmyAchievement)
		}
	}

	for _, player := range g.getAllPlayers() {
		for _, achievement := range player.achievements {
			if achievement.achievementType == LargestArmy {
				largestArmyAchievement = achievement
				player.achievements = slices.Remove(player.achievements, largestArmyAchievement)
			}
		}
	}

	if largestArmyAchievement == nil {
		return errors.WithStack(app_errors.ErrAchievementCardNotFound)
	}

	var playerHasLargestArmy *Player
	maxKnightDevelopmentCardPlayed := 0
	for _, player := range g.getAllPlayers() {
		knightDevelopmentCardPlayed := 0
		for _, developmentCard := range player.developmentCards {
			if developmentCard.developmentCardType == Knight && developmentCard.status == Used {
				knightDevelopmentCardPlayed++
			}
		}

		if knightDevelopmentCardPlayed == maxKnightDevelopmentCardPlayed {
			playerHasLargestArmy = nil
		}

		if knightDevelopmentCardPlayed >= 3 && knightDevelopmentCardPlayed > maxKnightDevelopmentCardPlayed {
			maxKnightDevelopmentCardPlayed = knightDevelopmentCardPlayed
			playerHasLargestArmy = player
		}
	}

	if playerHasLargestArmy != nil {
		playerHasLargestArmy.achievements = append(playerHasLargestArmy.achievements, largestArmyAchievement)
	} else {
		g.achievements = append(g.achievements, largestArmyAchievement)
	}

	return nil
}

func (g *Game) calculateScore() {
	for _, player := range g.getAllPlayers() {
		score := 0

		score += len(player.achievements) * 2

		for _, construction := range player.constructions {
			if construction.land != nil {
				switch construction.constructionType {
				case Settlement:
					score++
				case City:
					score += 2
				}
			}
		}

		for _, developmentCard := range player.developmentCards {
			if developmentCard.isVictoryPointCard() && developmentCard.status == Used {
				score++
			}
		}

		player.score = score

		if score >= 10 {
			g.status = Finished
		}
	}
}
