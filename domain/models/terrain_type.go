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
	Forest   terrainType = "FOREST"
	Hill     terrainType = "HILL"
	Field    terrainType = "FIELD"
	Pasture  terrainType = "PASTURE"
	Mountain terrainType = "MOUNTAIN"
	Desert   terrainType = "DESERT"
)
