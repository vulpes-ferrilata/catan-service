package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GameBuilder interface {
	SetID(id primitive.ObjectID) GameBuilder
	SetStatus(status gameStatus) GameBuilder
	SetTurn(turn int) GameBuilder
	SetIsRolledDices(isRolledDices bool) GameBuilder
	SetPlayers(players ...*Player) GameBuilder
	SetDices(dices ...*Dice) GameBuilder
	SetAchievements(achievements ...*Achievement) GameBuilder
	SetResourceCards(resourceCards ...*ResourceCard) GameBuilder
	SetDevelopmentCards(developmentCards ...*DevelopmentCard) GameBuilder
	SetTerrains(terrains ...*Terrain) GameBuilder
	SetHarbors(harbors ...*Harbor) GameBuilder
	SetRobber(robber *Robber) GameBuilder
	SetLands(lands ...*Land) GameBuilder
	SetPaths(paths ...*Path) GameBuilder
	SetVersion(version int) GameBuilder
	Create() *Game
}

func NewGameBuilder() GameBuilder {
	return &gameBuilder{}
}

type gameBuilder struct {
	id               primitive.ObjectID
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
	version          int
}

func (g *gameBuilder) SetID(id primitive.ObjectID) GameBuilder {
	g.id = id

	return g
}

func (g *gameBuilder) SetStatus(status gameStatus) GameBuilder {
	g.status = status

	return g
}

func (g *gameBuilder) SetTurn(turn int) GameBuilder {
	g.turn = turn

	return g
}

func (g *gameBuilder) SetIsRolledDices(isRolledDices bool) GameBuilder {
	g.isRolledDices = isRolledDices

	return g
}

func (g *gameBuilder) SetPlayers(players ...*Player) GameBuilder {
	g.players = players

	return g
}

func (g *gameBuilder) SetDices(dices ...*Dice) GameBuilder {
	g.dices = dices

	return g
}

func (g *gameBuilder) SetAchievements(achievements ...*Achievement) GameBuilder {
	g.achievements = achievements

	return g
}

func (g *gameBuilder) SetResourceCards(resourceCards ...*ResourceCard) GameBuilder {
	g.resourceCards = resourceCards

	return g
}

func (g *gameBuilder) SetDevelopmentCards(developmentCards ...*DevelopmentCard) GameBuilder {
	g.developmentCards = developmentCards

	return g
}

func (g *gameBuilder) SetTerrains(terrains ...*Terrain) GameBuilder {
	g.terrains = terrains

	return g
}

func (g *gameBuilder) SetHarbors(harbors ...*Harbor) GameBuilder {
	g.harbors = harbors

	return g
}

func (g *gameBuilder) SetRobber(robber *Robber) GameBuilder {
	g.robber = robber

	return g
}

func (g *gameBuilder) SetLands(lands ...*Land) GameBuilder {
	g.lands = lands

	return g
}

func (g *gameBuilder) SetPaths(paths ...*Path) GameBuilder {
	g.paths = paths

	return g
}

func (g *gameBuilder) SetVersion(version int) GameBuilder {
	g.version = version

	return g
}

func (g gameBuilder) Create() *Game {
	return &Game{
		aggregateRoot: aggregateRoot{
			aggregate: aggregate{
				id: g.id,
			},
			version: g.version,
		},
		status:           g.status,
		turn:             g.turn,
		isRolledDices:    g.isRolledDices,
		players:          g.players,
		dices:            g.dices,
		achievements:     g.achievements,
		resourceCards:    g.resourceCards,
		developmentCards: g.developmentCards,
		terrains:         g.terrains,
		harbors:          g.harbors,
		robber:           g.robber,
		lands:            g.lands,
		paths:            g.paths,
	}
}
