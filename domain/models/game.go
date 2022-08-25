package models

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	aggregateRoot
	status           gameStatus
	turn             int
	isRolledDices    bool
	players          []*Player
	dices            []*Dice
	achievements     []*Achievement
	resourceCards    []*ResourceCard
	developmentCards []*DevelopmentCard
	terrains         []*Terrain
	harbors          []*Harbor
	robber           *Robber
	lands            []*Land
	paths            []*Path
}

func (g Game) GetStatus() gameStatus {
	return g.status
}

func (g Game) GetTurn() int {
	return g.turn
}

func (g Game) IsRolledDices() bool {
	return g.isRolledDices
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

func (g Game) GetHarbors() []*Harbor {
	return slices.Clone(g.harbors)
}

func (g Game) GetRobber() *Robber {
	return g.robber
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

func (g *Game) NewPlayer(userID primitive.ObjectID) error {
	state := g.getState()

	if err := state.newPlayer(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) Start(userID primitive.ObjectID) error {
	state := g.getState()

	if err := state.startGame(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) BuildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	state := g.getState()

	if err := state.buildSettlementAndRoad(userID, landID, pathID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) RollDices(userID primitive.ObjectID) error {
	state := g.getState()

	if err := state.rollDices(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) MoveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	state := g.getState()

	if err := state.moveRobber(userID, terrainID, playerID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Game) EndTurn(userID primitive.ObjectID) error {
	state := g.getState()

	if err := state.endTurn(userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
