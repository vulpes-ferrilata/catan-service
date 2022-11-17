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
	LumberHarbor  = HarborType{"Lumber"}
	BrickHarbor   = HarborType{"Brick"}
	WoolHarbor    = HarborType{"Wool"}
	GrainHarbor   = HarborType{"Grain"}
	OreHarbor     = HarborType{"Ore"}
	GeneralHarbor = HarborType{"General"}
)
