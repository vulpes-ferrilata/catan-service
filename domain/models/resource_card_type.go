package models

import (
	"github.com/pkg/errors"
)

func NewResourceCardType(enum string) (ResourceCardType, error) {
	resourceCardType := ResourceCardType{enum}

	switch resourceCardType {
	case Lumber, Brick, Wool, Grain, Ore:
		return resourceCardType, nil
	default:
		return resourceCardType, errors.New("resource type is invalid")
	}
}

type ResourceCardType struct {
	enum string
}

func (r ResourceCardType) String() string {
	return r.enum
}

var (
	Lumber = ResourceCardType{"Lumber"}
	Brick  = ResourceCardType{"Brick"}
	Wool   = ResourceCardType{"Wool"}
	Grain  = ResourceCardType{"Grain"}
	Ore    = ResourceCardType{"Ore"}
)
