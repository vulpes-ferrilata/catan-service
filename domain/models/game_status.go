package models

import "github.com/pkg/errors"

func NewGameStatus(value string) (GameStatus, error) {
	gameStatus := GameStatus(value)

	switch gameStatus {
	case Waiting, Started, Finished:
		return gameStatus, nil
	default:
		return gameStatus, errors.Wrap(ErrEnumIsInvalid, string(gameStatus))
	}
}

type GameStatus string

const (
	Waiting  GameStatus = "Waiting"
	Started  GameStatus = "Started"
	Finished GameStatus = "Finished"
)
