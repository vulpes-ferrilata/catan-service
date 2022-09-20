package models

import (
	"github.com/pkg/errors"
)

func NewHexCornerLocation(enum string) (HexCornerLocation, error) {
	hexCornerLocation := HexCornerLocation{enum}
	switch hexCornerLocation {
	case Top, Bottom:
		return hexCornerLocation, nil
	default:
		return hexCornerLocation, errors.New("hex corner location is invalid")
	}
}

type HexCornerLocation struct {
	enum string
}

func (h HexCornerLocation) String() string {
	return h.enum
}

var (
	Top    = HexCornerLocation{"TOP"}
	Bottom = HexCornerLocation{"BOTTOM"}
)
