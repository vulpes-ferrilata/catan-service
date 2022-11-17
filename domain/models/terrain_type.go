package models

import (
	"github.com/pkg/errors"
)

func NewTerrainType(value string) (terrainType, error) {
	terrainType := terrainType(value)
	switch terrainType {
	case Hill, Field, Pasture, Mountain, Forest, Desert:
		return terrainType, nil
	default:
		return terrainType, errors.New("terrain is invalid")
	}
}

type terrainType string

const (
	Forest   terrainType = "Forest"
	Hill     terrainType = "Hill"
	Field    terrainType = "Field"
	Pasture  terrainType = "Pasture"
	Mountain terrainType = "Mountain"
	Desert   terrainType = "Desert"
)
