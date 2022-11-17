package models

import (
	"github.com/pkg/errors"
)

func NewGameStatus(enum string) (GameStatus, error) {
	gameStatus := GameStatus{enum}

	switch gameStatus {
	case Waiting, Started, Finished:
		return gameStatus, nil
	default:
		return gameStatus, errors.Errorf("game status is invalid")
	}
}

type GameStatus struct {
	enum string
}

func (g GameStatus) String() string {
	return g.enum
}

var (
	Waiting  = GameStatus{"Waiting"}
	Started  = GameStatus{"Started"}
	Finished = GameStatus{"Finished"}
)
