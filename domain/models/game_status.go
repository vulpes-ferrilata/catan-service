package models

import (
	"github.com/pkg/errors"
)

func NewGameStatus(value string) (gameStatus, error) {
	gameStatus := gameStatus(value)

	switch gameStatus {
	case Waiting, Started, Finished:
		return gameStatus, nil
	default:
		return gameStatus, errors.Errorf("game status is invalid")
	}
}

type gameStatus string

const (
	Waiting  gameStatus = "WAITING"
	Started  gameStatus = "STARTED"
	Finished gameStatus = "FINISHED"
)
