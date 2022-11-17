package models

import (
	"github.com/pkg/errors"
)

func NewGamePhase(enum string) (GamePhase, error) {
	gamePhase := GamePhase{enum}

	switch gamePhase {
	case Setup, ResourceProduction, ResourceDiscard, Robbing, ResourceConsumption:
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
	Setup               = GamePhase{"Setup"}
	ResourceProduction  = GamePhase{"ResourceProduction"}
	ResourceDiscard     = GamePhase{"ResourceDiscard"}
	Robbing             = GamePhase{"Robbing"}
	ResourceConsumption = GamePhase{"ResourceConsumption"}
)
