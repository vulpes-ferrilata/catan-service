package models

import (
	"math/rand"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
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
	return slices.Clone(g.players)
}

func (g Game) GetDices() []*Dice {
	return slices.Clone(g.dices)
}

func (g Game) GetAchievements() []*Achievement {
	return slices.Clone(g.achievements)
}

func (g Game) GetResourceCards() []*ResourceCard {
	return slices.Clone(g.resourceCards)
}

func (g Game) GetDevelopmentCards() []*DevelopmentCard {
	return slices.Clone(g.developmentCards)
}

func (g Game) GetTerrains() []*Terrain {
	return slices.Clone(g.terrains)
}

func (g Game) GetLands() []*Land {
	return slices.Clone(g.lands)
}

func (g Game) GetPaths() []*Path {
	return slices.Clone(g.paths)
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

	return nil
}

func (g *Game) RollDices(userID primitive.ObjectID) error {
	if err := g.getState().rollDices(userID); err != nil {
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

	return nil
}

func (g *Game) BuildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	if err := g.getState().buildRoad(userID, pathID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) UpgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	if err := g.getState().upgradeCity(userID, constructionID); err != nil {
		return errors.WithStack(err)
	}

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

func (g *Game) MaritimeTrade(userID primitive.ObjectID, demandingResourceCardType ResourceCardType) error {
	if err := g.getState().maritimeTrade(userID, demandingResourceCardType); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) OfferTrading(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	if err := g.getState().offerTrading(userID, playerID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) ConfirmTrading(userID primitive.ObjectID) error {
	if err := g.getState().confirmTrading(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) CancelTrading(userID primitive.ObjectID) error {
	if err := g.getState().cancelTrading(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) PlayKnightCard(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	if err := g.getState().playKnightCard(userID, terrainID, playerID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) PlayRoadBuildingCard(userID primitive.ObjectID, pathIDs []primitive.ObjectID) error {
	if err := g.getState().playRoadBuildingCard(userID, pathIDs); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) PlayYearOfPlentyCard(userID primitive.ObjectID, resourceCardTypes []ResourceCardType) error {
	if err := g.getState().playYearOfPlentyCard(userID, resourceCardTypes); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) PlayMonopolyCard(userID primitive.ObjectID, resourceCardType ResourceCardType) error {
	if err := g.getState().playMonopolyCard(userID, resourceCardType); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) useResourceCards(resourceCardTypes ...ResourceCardType) error {
	for _, resourceCardType := range resourceCardTypes {
		resourceCard, isExists := slices.Find(func(resourceCard *ResourceCard) bool {
			return resourceCard.resourceCardType == resourceCardType
		}, g.activePlayer.resourceCards)
		if !isExists {
			return errors.WithStack(app_errors.ErrYouHaveInsufficientResourceCards)
		}

		g.activePlayer.resourceCards = slices.Remove(g.activePlayer.resourceCards, resourceCard)
		g.resourceCards = append(g.resourceCards, resourceCard)
	}

	return nil
}

func (g *Game) buildSettlement(land *Land) error {
	isLandAdjacentToAnyConstruction := slices.Any(func(player *Player) bool {
		return slices.Any(func(construction *Construction) bool {
			return construction.land != nil && construction.land.hexCorner.isAdjacentWithHexCorner(land.hexCorner)
		}, player.constructions)
	}, g.players)
	if isLandAdjacentToAnyConstruction {
		return errors.WithStack(app_errors.ErrNearbyLandsMustBeVacant)
	}

	isLandAdjacentToPlayerRoad := slices.Any(func(road *Road) bool {
		return road.path != nil && road.path.hexEdge.isAdjacentWithHexCorner(land.hexCorner)
	}, g.activePlayer.roads)
	if !isLandAdjacentToPlayerRoad {
		return errors.WithStack(app_errors.ErrSelectedLandMustBeAdjacentToYourRoad)
	}

	settlement, isExists := slices.Find(func(construction *Construction) bool {
		return construction.constructionType == Settlement && construction.land == nil
	}, g.activePlayer.constructions)
	if !isExists {
		return errors.WithStack(app_errors.ErrYouRunOutOfSettlements)
	}

	g.lands = slices.Remove(g.lands, land)
	settlement.land = land

	return nil
}

func (g *Game) buildRoad(path *Path) error {
	isAdjacentToConstruction := slices.Any(func(construction *Construction) bool {
		return construction.land != nil && construction.land.hexCorner.isAdjacentWithHexEdge(path.hexEdge)
	}, g.activePlayer.constructions)

	//if selected path adjacent to your construction, other adjacent land cannot be occupied by other player
	if !isAdjacentToConstruction {
		adjacentHexCorners := findAdjacentHexCornersFromHexEdge(path.hexEdge)

		intersectionHexCorners := slices.Filter(func(adjacentHexCorner HexCorner) bool {
			return slices.Any(func(road *Road) bool {
				return road.path != nil && adjacentHexCorner.isAdjacentWithHexEdge(road.path.hexEdge)
			}, g.activePlayer.roads)
		}, adjacentHexCorners)

		if len(intersectionHexCorners) == 0 {
			return errors.WithStack(app_errors.ErrSelectedPathMustBeAdjacentToYourConstructionOrRoad)
		}

		if len(intersectionHexCorners) == 1 {
			isSelectedPathPassThroughConstructionOfOtherPlayer := slices.Any(func(otherPlayer *Player) bool {
				return slices.Any(func(construction *Construction) bool {
					return construction.land != nil && construction.land.hexCorner == intersectionHexCorners[0]
				}, otherPlayer.constructions)
			}, g.players)
			if isSelectedPathPassThroughConstructionOfOtherPlayer {
				return errors.WithStack(app_errors.ErrSelectedPathPassThroughConstructionOfOtherPlayer)
			}
		}

		//if you have roads on both sides of selected path, both sides cannot be occupied
	}

	road, isExists := slices.Find(func(road *Road) bool {
		return road.path == nil
	}, g.activePlayer.roads)
	if !isExists {
		return errors.WithStack(app_errors.ErrYouRunOutOfRoads)
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

	city, isExists := slices.Find(func(construction *Construction) bool {
		return construction.land == nil && construction.constructionType == City
	}, g.activePlayer.constructions)
	if !isExists {
		return errors.WithStack(app_errors.ErrYouRunOutOfCities)
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
	terrainHasRobber, isExists := slices.Find(func(terrain *Terrain) bool {
		return terrain.robber != nil
	}, g.terrains)
	if !isExists {
		return errors.WithStack(app_errors.ErrTerrainNotFound)
	}

	if player == nil {
		hasPlayerCanBeRob := slices.Any(func(player *Player) bool {
			return slices.Any(func(construction *Construction) bool {
				return construction.land != nil && construction.land.hexCorner.isAdjacentWithHex(terrainHasRobber.hex)
			}, player.constructions)
		}, g.players)
		if hasPlayerCanBeRob {
			return errors.WithStack(app_errors.ErrYouMustRobPlayerWhoHasConstructionNextToRobber)
		}

		return nil
	}

	canBeRob := slices.Any(func(construction *Construction) bool {
		return construction.land != nil && construction.land.hexCorner.isAdjacentWithHex(terrainHasRobber.hex)
	}, player.constructions)
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

func (g *Game) useDevelopmentCard(developmentCardType DevelopmentCardType) error {
	knightDevelopmentCard, isExists := slices.Find(func(developmentCard *DevelopmentCard) bool {
		return developmentCard.developmentCardType == developmentCardType && developmentCard.status == Enable
	}, g.activePlayer.developmentCards)
	if !isExists {
		return errors.WithStack(app_errors.ErrDevelopmentCardNotFound)
	}

	knightDevelopmentCard.status = Used

	for _, developmentCard := range g.activePlayer.developmentCards {
		if developmentCard.status == Enable {
			developmentCard.status = Disable
		}
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

	if longestRoadAchievement == nil {
		return errors.WithStack(app_errors.ErrAchievementCardNotFound)
	}

	var playerHasLongestRoad *Player
	maxLongestRoad := 0
	for _, player := range g.getAllPlayers() {
		longestRoad := g.calculateLongestRoad(player)

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

func (g Game) calculateLongestRoad(player *Player) int {
	longestRoad := 0

	for _, road := range player.roads {
		remainRoads := slices.Remove(player.roads, road)
		totalRoads := g.calculateLongestRoadFromCurrentRoad(player, road, remainRoads)
		if totalRoads > longestRoad {
			longestRoad = totalRoads
		}
	}

	return longestRoad
}

func (g Game) calculateLongestRoadFromCurrentRoad(player *Player, currentRoad *Road, otherRoads []*Road) int {
	longestRoad := 0

	if currentRoad.path == nil {
		return 0
	}

	for _, otherRoad := range otherRoads {
		if otherRoad.path == nil {
			continue
		}

		hexCorner, isExists := findIntersectionHexCornerBetweenTwoHexEdges(currentRoad.path.hexEdge, otherRoad.path.hexEdge)
		if !isExists {
			continue
		}

		otherPlayers := slices.Remove(g.players, player)

		isOtherPlayerHasConstructionBetweenTwoRoads := slices.Any(func(otherPlayer *Player) bool {
			return slices.Any(func(construction *Construction) bool {
				return construction.land != nil && construction.land.hexCorner == hexCorner
			}, otherPlayer.constructions)
		}, otherPlayers)
		if isOtherPlayerHasConstructionBetweenTwoRoads {
			continue
		}

		remainRoads := slices.Remove(otherRoads, otherRoad)
		result := g.calculateLongestRoadFromCurrentRoad(player, otherRoad, remainRoads)
		if result > longestRoad {
			longestRoad = result
		}
	}

	return 1 + longestRoad
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
	maxScore := 0

	for _, player := range g.getAllPlayers() {
		score := 0

		score += len(player.achievements) * 2

		for _, construction := range player.constructions {
			if construction.land != nil {
				switch construction.constructionType {
				case Settlement:
					score += 1
				case City:
					score += 2
				}
			}
		}

		for _, developmentCard := range player.developmentCards {
			switch developmentCard.developmentCardType {
			case VictoryPoints:
				score += 1
			}
		}

		player.score = score

		if score > maxScore {
			maxScore = score
		}
	}

	if maxScore >= 10 {
		g.status = Finished
	}
}
