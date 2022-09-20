package models

import (
	"github.com/pkg/errors"
)

func NewHexEdgeLocation(enum string) (HexEdgeLocation, error) {
	hexEdgeLocation := HexEdgeLocation{enum}
	switch hexEdgeLocation {
	case TopLeft, MiddleLeft, BottomLeft:
		return hexEdgeLocation, nil
	default:
		return hexEdgeLocation, errors.New("hex edge location is invalid")
	}
}

type HexEdgeLocation struct {
	enum string
}

func (h HexEdgeLocation) String() string {
	return h.enum
}

var (
	TopLeft    = HexEdgeLocation{"TOP_LEFT"}
	MiddleLeft = HexEdgeLocation{"MIDDLE_LEFT"}
	BottomLeft = HexEdgeLocation{"BOTTOM_LEFT"}
)
