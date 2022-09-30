package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GameBuilder struct {
	id               primitive.ObjectID
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
	version          int
}

func (g GameBuilder) SetID(id primitive.ObjectID) GameBuilder {
	g.id = id

	return g
}

func (g GameBuilder) SetStatus(status GameStatus) GameBuilder {
	g.status = status

	return g
}

func (g GameBuilder) SetPhase(phase GamePhase) GameBuilder {
	g.phase = phase

	return g
}

func (g GameBuilder) SetTurn(turn int) GameBuilder {
	g.turn = turn

	return g
}

func (g GameBuilder) SetActivePlayer(activePlayer *Player) GameBuilder {
	g.activePlayer = activePlayer

	return g
}

func (g GameBuilder) SetPlayers(players ...*Player) GameBuilder {
	g.players = players

	return g
}

func (g GameBuilder) SetDices(dices ...*Dice) GameBuilder {
	g.dices = dices

	return g
}

func (g GameBuilder) SetAchievements(achievements ...*Achievement) GameBuilder {
	g.achievements = achievements

	return g
}

func (g GameBuilder) SetResourceCards(resourceCards ...*ResourceCard) GameBuilder {
	g.resourceCards = resourceCards

	return g
}

func (g GameBuilder) SetDevelopmentCards(developmentCards ...*DevelopmentCard) GameBuilder {
	g.developmentCards = developmentCards

	return g
}

func (g GameBuilder) SetTerrains(terrains ...*Terrain) GameBuilder {
	g.terrains = terrains

	return g
}

func (g GameBuilder) SetLands(lands ...*Land) GameBuilder {
	g.lands = lands

	return g
}

func (g GameBuilder) SetPaths(paths ...*Path) GameBuilder {
	g.paths = paths

	return g
}

func (g GameBuilder) SetVersion(version int) GameBuilder {
	g.version = version

	return g
}

func (g GameBuilder) Create() *Game {
	return &Game{
		aggregateRoot: aggregateRoot{
			aggregate: aggregate{
				id: g.id,
			},
			version: g.version,
		},
		status:           g.status,
		phase:            g.phase,
		turn:             g.turn,
		activePlayer:     g.activePlayer,
		players:          g.players,
		dices:            g.dices,
		achievements:     g.achievements,
		resourceCards:    g.resourceCards,
		developmentCards: g.developmentCards,
		terrains:         g.terrains,
		lands:            g.lands,
		paths:            g.paths,
	}
}
