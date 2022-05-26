package models

import (
	service "github.com/VulpesFerrilata/catan-service"
	"github.com/VulpesFerrilata/catan-service/domain/models/common"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewGame(id uuid.UUID,
	status GameStatus,
	activePlayerID uuid.UUID,
	turn int,
	players []*Player,
	dices []*Dice,
	achievements []*Achievement) *Game {
	return &Game{
		Entity:       common.NewEntity(id),
		status:       status,
		turn:         turn,
		players:      players,
		dices:        dices,
		achievements: achievements,
	}
}

type Game struct {
	common.Entity
	status         GameStatus
	activePlayerID uuid.UUID
	turn           int
	players        []*Player
	dices          []*Dice
	achievements   []*Achievement
}

func (g Game) GetStatus() GameStatus {
	return g.status
}

func (g Game) GetActivePlayerID() uuid.UUID {
	return g.activePlayerID
}

func (g Game) GetTurn() int {
	return g.turn
}

func (g Game) GetPlayers() []*Player {
	return g.players
}

func (g Game) GetDices() []*Dice {
	return g.dices
}

func (g *Game) AddPlayer(player *Player) error {
	if len(g.players) >= 4 {
		return errors.WithStack(service.ErrPlayerIsFull)
	}

	for _, currentPlayer := range g.players {
		if currentPlayer.GetUserID() == player.GetUserID() {
			return errors.WithStack(service.ErrPlayerAlreadyJoined)
		}
	}

	g.players = append(g.players, player)

	return nil
}

func (g *Game) Start(userID uuid.UUID) error {
	if g.activePlayerID != userID {
		return errors.WithStack(service.ErrOtherPlayerTurn)
	}

	g.status = Started

	return nil
}
