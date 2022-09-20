package models

import (
	"github.com/pkg/errors"
)

func NewGamePhase(enum string) (GamePhase, error) {
	gamePhase := GamePhase{enum}

	switch gamePhase {
	case ResourceProduction, Robbing, ResourceConsumption:
		return gamePhase, nil
	default:
		return gamePhase, errors.Errorf("game phase is invalid")
	}
}

type GamePhase struct {
	enum string
}

func (g GamePhase) String() string {
	return g.enum
}

var (
	ResourceProduction  = GamePhase{"RESOURCE_PRODUCTION"}
	Robbing             = GamePhase{"ROBBING"}
	ResourceConsumption = GamePhase{"RESOURCE_CONSUMPTION"}
)
