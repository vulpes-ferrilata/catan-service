package models

import (
	"github.com/pkg/errors"
)

func NewHarborType(value string) (harborType, error) {
	harborType := harborType(value)

	switch harborType {
	case LumberHarbor, BrickHarbor, WoolHarbor, GrainHarbor, OreHarbor, GeneralHarbor:
		return harborType, nil
	default:
		return harborType, errors.New("harbor type is invalid")
	}
}

type harborType string

const (
	LumberHarbor  harborType = "LUMBER"
	BrickHarbor   harborType = "BRICK"
	WoolHarbor    harborType = "WOOL"
	GrainHarbor   harborType = "GRAIN"
	OreHarbor     harborType = "ORE"
	GeneralHarbor harborType = "GENERAL"
)
