package models

import (
	"github.com/pkg/errors"
)

func NewHexEdgeLocation(value string) (hexEdgeLocation, error) {
	hexEdgeLocation := hexEdgeLocation(value)
	switch hexEdgeLocation {
	case TopLeft, MiddleLeft, BottomLeft:
		return hexEdgeLocation, nil
	default:
		return hexEdgeLocation, errors.New("hex edge location is invalid")
	}
}

type hexEdgeLocation string

const (
	TopLeft    hexEdgeLocation = "TOP_LEFT"
	MiddleLeft hexEdgeLocation = "MIDDLE_LEFT"
	BottomLeft hexEdgeLocation = "BOTTOM_LEFT"
)
