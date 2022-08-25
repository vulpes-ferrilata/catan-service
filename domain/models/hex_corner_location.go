package models

import (
	"github.com/pkg/errors"
)

func NewHexCornerLocation(value string) (hexCornerLocation, error) {
	hexCornerLocation := hexCornerLocation(value)
	switch hexCornerLocation {
	case Top, Bottom:
		return hexCornerLocation, nil
	default:
		return hexCornerLocation, errors.New("hex corner location is invalid")
	}
}

type hexCornerLocation string

const (
	Top    hexCornerLocation = "TOP"
	Bottom hexCornerLocation = "BOTTOM"
)
