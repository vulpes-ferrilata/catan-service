package models

import (
	"github.com/pkg/errors"
)

func NewPlayerColor(enum string) (PlayerColor, error) {
	color := PlayerColor{enum}

	switch color {
	case Red, Blue, Green, Yellow:
		return color, nil
	default:
		return color, errors.New("player color is invalid")
	}
}

type PlayerColor struct {
	enum string
}

func (p PlayerColor) String() string {
	return p.enum
}

var (
	Red    = PlayerColor{"RED"}
	Blue   = PlayerColor{"BLUE"}
	Green  = PlayerColor{"GREEN"}
	Yellow = PlayerColor{"YELLOW"}
)
