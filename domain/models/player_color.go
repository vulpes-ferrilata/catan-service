package models

import (
	"github.com/pkg/errors"
)

func NewPlayerColor(value string) (playerColor, error) {
	color := playerColor(value)

	switch color {
	case Red, Blue, Green, Yellow:
		return color, nil
	default:
		return color, errors.New("color is invalid")
	}
}

type playerColor string

const (
	Red    playerColor = "RED"
	Blue   playerColor = "BLUE"
	Green  playerColor = "GREEN"
	Yellow playerColor = "YELLOW"
)
