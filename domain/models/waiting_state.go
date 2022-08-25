package models

import (
	"math/rand"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type waitingState struct {
	game *Game
}

func (w waitingState) newPlayer(userID primitive.ObjectID) error {
	if len(w.game.players) >= 4 {
		return errors.WithStack(app_errors.ErrGameHasFullPlayers)
	}

	for _, player := range w.game.players {
		if player.userID == userID {
			return errors.WithStack(app_errors.ErrYouAlreadyJoined)
		}
	}

	colors := []playerColor{Red, Blue, Green, Yellow}
	turnOrder := 1
	for _, player := range w.game.players {
		colors = slices.Remove(colors, player.color)
		if player.turnOrder >= turnOrder {
			turnOrder = player.turnOrder + 1
		}
	}

	isActive := len(w.game.players) == 0

	id := primitive.NewObjectID()

	player := NewPlayerBuilder().
		SetID(id).
		SetUserID(userID).
		SetColor(colors[0]).
		SetTurnOrder(turnOrder).
		SetIsActive(isActive).
		Create()

	w.game.players = append(w.game.players, player)

	return nil
}

func (w waitingState) initDices() {
	for i := 1; i <= 2; i++ {
		dice := &Dice{
			aggregate: aggregate{
				id: primitive.NewObjectID(),
			},
			number: 1,
		}
		w.game.dices = append(w.game.dices, dice)
	}
}

func (w waitingState) initAchievements() {
	longestRoadAchievement := &Achievement{
		aggregate: aggregate{
			id: primitive.NewObjectID(),
		},
		achievementType: LongestRoad,
	}
	w.game.achievements = append(w.game.achievements, longestRoadAchievement)

	largestArmyAchievement := &Achievement{
		aggregate: aggregate{
			id: primitive.NewObjectID(),
		},
		achievementType: LargestArmy,
	}
	w.game.achievements = append(w.game.achievements, largestArmyAchievement)
}

func (w waitingState) initResourceCards() {
	resourceCardTypes := []resourceCardType{
		Lumber,
		Brick,
		Wool,
		Grain,
		Ore,
	}

	for _, resourceCardType := range resourceCardTypes {
		for i := 1; i <= 19; i++ {
			resourceCard := &ResourceCard{
				aggregate: aggregate{
					id: primitive.NewObjectID(),
				},
				resourceCardType: resourceCardType,
				isSelected:       false,
			}
			w.game.resourceCards = append(w.game.resourceCards, resourceCard)
		}
	}
}

func (w waitingState) initDevelopmentCards() {
	for i := 1; i <= 14; i++ {
		developmentCard := &DevelopmentCard{
			aggregate: aggregate{
				id: primitive.NewObjectID(),
			},
			developmentCardType: Knight,
			status:              Disable,
		}
		w.game.developmentCards = append(w.game.developmentCards, developmentCard)
	}

	progressDevelopmentCardTypes := []developmentCardType{
		Monopoly,
		RoadBuilding,
		YearOfPlenty,
	}
	for _, progressDevelopmentCardType := range progressDevelopmentCardTypes {
		for i := 1; i <= 2; i++ {
			developmentCard := &DevelopmentCard{
				aggregate: aggregate{
					id: primitive.NewObjectID(),
				},
				developmentCardType: progressDevelopmentCardType,
				status:              Disable,
			}
			w.game.developmentCards = append(w.game.developmentCards, developmentCard)
		}
	}

	for i := 1; i <= 5; i++ {
		developmentCard := &DevelopmentCard{
			aggregate: aggregate{
				id: primitive.NewObjectID(),
			},
			developmentCardType: VictoryPoints,
			status:              Disable,
		}
		w.game.developmentCards = append(w.game.developmentCards, developmentCard)
	}
}

func (w waitingState) initTerrains() {
	spiralHexes := make([]Hex, 0)

	hexDirections := []hexDirection{
		NewHexDirection(1, 0),
		NewHexDirection(0, 1),
		NewHexDirection(-1, 1),
		NewHexDirection(-1, 0),
		NewHexDirection(0, -1),
		NewHexDirection(1, -1),
	}

	//reverse
	reverse := rand.Intn(2) != 0
	if reverse {
		for i, j := 0, len(hexDirections)-1; i < j; i, j = i+1, j-1 {
			hexDirections[i], hexDirections[j] = hexDirections[j], hexDirections[i]
		}
	}

	//shuffle
	hexDirectionIdx := rand.Intn(len(hexDirections))
	hexDirections = append(hexDirections[hexDirectionIdx:], hexDirections[:hexDirectionIdx]...)

	centerHex := NewHex(0, 0)
	spiralHexes = append(spiralHexes, centerHex)

	for radius := 1; radius <= 2; radius++ {
		circleHexes := make([]Hex, 0)
		hex := hexDirections[len(hexDirections)-2].Multiply(radius).CalculateEndpoint(centerHex)
		for _, hexDirection := range hexDirections {
			for step := 0; step < radius; step++ {
				hex = hexDirection.CalculateEndpoint(hex)
				circleHexes = append(circleHexes, hex)
			}
		}
		spiralHexes = append(spiralHexes, circleHexes...)
	}

	numbers := []int{11, 3, 6, 5, 4, 9, 10, 8, 4, 11, 12, 9, 10, 8, 3, 6, 2, 5}

	terrainTypes := []terrainType{
		Field,
		Field,
		Field,
		Field,
		Forest,
		Forest,
		Forest,
		Forest,
		Pasture,
		Pasture,
		Pasture,
		Pasture,
		Mountain,
		Mountain,
		Mountain,
		Hill,
		Hill,
		Hill,
	}
	rand.Shuffle(len(terrainTypes), func(i, j int) { terrainTypes[i], terrainTypes[j] = terrainTypes[j], terrainTypes[i] })

	desertIdx := rand.Intn(len(spiralHexes))

	if desertIdx == len(spiralHexes)-1 {
		numbers = append(numbers, 7)
		terrainTypes = append(terrainTypes, Desert)
	} else {
		numbers = append(numbers[:desertIdx+1], numbers[desertIdx:]...)
		numbers[desertIdx] = 7

		terrainTypes = append(terrainTypes[:desertIdx+1], terrainTypes[desertIdx:]...)
		terrainTypes[desertIdx] = Desert
	}

	for i := len(spiralHexes) - 1; i >= 0; i-- {
		terrain := &Terrain{
			aggregate: aggregate{
				id: primitive.NewObjectID(),
			},
			hex:         spiralHexes[i],
			number:      numbers[i],
			terrainType: terrainTypes[i],
		}
		w.game.terrains = append(w.game.terrains, terrain)
	}
}

func (w waitingState) initHarbors() {
	circleHexes := make([]Hex, 0)

	hexDirections := []hexDirection{
		NewHexDirection(1, 0),
		NewHexDirection(0, 1),
		NewHexDirection(-1, 1),
		NewHexDirection(-1, 0),
		NewHexDirection(0, -1),
		NewHexDirection(1, -1),
	}

	centerHex := NewHex(0, 0)

	hex := hexDirections[len(hexDirections)-2].Multiply(3).CalculateEndpoint(centerHex)
	for _, hexDirection := range hexDirections {
		for step := 0; step < 3; step++ {
			hex = hexDirection.CalculateEndpoint(hex)
			circleHexes = append(circleHexes, hex)
		}
	}

	oddOrEven := rand.Intn(2)
	oddOrEvenCircleHexes := make([]Hex, 0)
	for idx, circleHex := range circleHexes {
		if idx%2 == oddOrEven {
			oddOrEvenCircleHexes = append(oddOrEvenCircleHexes, circleHex)
		}
	}

	harborTypes := []harborType{
		WoolHarbor,
		LumberHarbor,
		BrickHarbor,
		OreHarbor,
		GrainHarbor,
		GeneralHarbor,
		GeneralHarbor,
		GeneralHarbor,
		GeneralHarbor,
	}
	rand.Shuffle(len(harborTypes), func(i, j int) { harborTypes[i], harborTypes[j] = harborTypes[j], harborTypes[i] })

	for idx, hex := range oddOrEvenCircleHexes {
		for _, terrain := range w.game.terrains {
			if terrain.hex.IsAdjacentWithHex(hex) {
				harbor := &Harbor{
					aggregate: aggregate{
						id: primitive.NewObjectID(),
					},
					terrainID:  terrain.id,
					hex:        hex,
					harborType: harborTypes[idx],
				}
				w.game.harbors = append(w.game.harbors, harbor)
				break
			}
		}
	}
}

func (w waitingState) initRobber() {
	for _, terrain := range w.game.terrains {
		if terrain.terrainType == Desert {
			w.game.robber = &Robber{
				aggregate: aggregate{
					id: primitive.NewObjectID(),
				},
				terrainID: terrain.id,
				isMoving:  false,
			}
		}
	}
}

func (w waitingState) initPaths() {
	hexEdges := make(map[HexEdge]struct{}, 0)
	for _, terrain := range w.game.terrains {
		adjacentHexEdges := FindAdjacentHexEdges(terrain.hex)
		for _, adjacentHexEdge := range adjacentHexEdges {
			hexEdges[adjacentHexEdge] = struct{}{}
		}
	}

	for hexEdge := range hexEdges {
		path := &Path{
			aggregate: aggregate{
				id: primitive.NewObjectID(),
			},
			hexEdge: hexEdge,
		}
		w.game.paths = append(w.game.paths, path)
	}
}

func (w waitingState) initLands() {
	hexCorners := make(map[HexCorner]struct{})
	for _, terrain := range w.game.terrains {
		adjacentHexCorners := FindAdjacentHexCorners(terrain.hex)
		for _, adjacentHexCorner := range adjacentHexCorners {
			hexCorners[adjacentHexCorner] = struct{}{}
		}
	}

	for hexCorner := range hexCorners {
		land := &Land{
			aggregate: aggregate{
				id: primitive.NewObjectID(),
			},
			hexCorner: hexCorner,
		}
		w.game.lands = append(w.game.lands, land)
	}
}

func (w waitingState) initConstructions() {
	for _, player := range w.game.players {
		constructions := make([]*Construction, 0)

		for i := 1; i <= 5; i++ {
			construction := &Construction{
				aggregate: aggregate{
					id: primitive.NewObjectID(),
				},
				constructionType: Settlement,
				land:             nil,
			}
			constructions = append(constructions, construction)
		}

		for i := 1; i <= 4; i++ {
			construction := &Construction{
				aggregate: aggregate{
					id: primitive.NewObjectID(),
				},
				constructionType: City,
				land:             nil,
			}
			constructions = append(constructions, construction)
		}

		player.constructions = constructions
	}
}

func (w waitingState) initRoads() {
	for _, player := range w.game.players {
		roads := make([]*Road, 0)

		for i := 1; i <= 15; i++ {
			road := &Road{
				aggregate: aggregate{
					id: primitive.NewObjectID(),
				},
				path: nil,
			}
			roads = append(roads, road)
		}

		player.roads = roads
	}
}

func (w waitingState) startGame(userID primitive.ObjectID) error {
	activePlayer, isExists := slices.Find(func(player *Player) bool {
		return player.userID == userID
	}, w.game.players)
	if !isExists {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if !activePlayer.isActive {
		return errors.WithStack(app_errors.ErrYouAreNotInTurn)
	}

	w.initDices()
	w.initAchievements()
	w.initResourceCards()
	w.initDevelopmentCards()
	w.initTerrains()
	w.initHarbors()
	w.initRobber()
	w.initPaths()
	w.initLands()
	w.initConstructions()
	w.initRoads()

	rand.Shuffle(len(w.game.players), func(i, j int) {
		w.game.players[i].turnOrder, w.game.players[j].turnOrder = w.game.players[j].turnOrder, w.game.players[i].turnOrder
	})

	for _, player := range w.game.players {
		if player.turnOrder == 1 {
			player.isActive = true
		} else {
			player.isActive = false
		}
	}

	w.game.status = Started

	return nil
}

func (w waitingState) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameHasNotStartedYet)
}

func (w waitingState) rollDices(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameHasNotStartedYet)
}

func (w waitingState) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameHasNotStartedYet)
}

func (w waitingState) endTurn(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrGameHasNotStartedYet)
}
