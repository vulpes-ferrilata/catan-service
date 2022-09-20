package models

import (
	"github.com/pkg/errors"
)

func NewHarborType(enum string) (HarborType, error) {
	harborType := HarborType{enum}

	switch harborType {
	case LumberHarbor, BrickHarbor, WoolHarbor, GrainHarbor, OreHarbor, GeneralHarbor:
		return harborType, nil
	default:
		return harborType, errors.New("harbor type is invalid")
	}
}

type HarborType struct {
	enum string
}

func (h HarborType) String() string {
	return h.enum
}

var (
	LumberHarbor  = HarborType{"LUMBER"}
	BrickHarbor   = HarborType{"BRICK"}
	WoolHarbor    = HarborType{"WOOL"}
	GrainHarbor   = HarborType{"GRAIN"}
	OreHarbor     = HarborType{"ORE"}
	GeneralHarbor = HarborType{"GENERAL"}
)
