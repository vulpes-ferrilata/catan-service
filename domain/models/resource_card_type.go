package models

import (
	"github.com/pkg/errors"
)

func NewResourceCardType(value string) (resourceCardType, error) {
	resourceCardType := resourceCardType(value)

	switch resourceCardType {
	case Lumber, Brick, Wool, Grain, Ore:
		return resourceCardType, nil
	default:
		return resourceCardType, errors.New("resource type is invalid")
	}
}

type resourceCardType string

const (
	Lumber resourceCardType = "LUMBER"
	Brick  resourceCardType = "BRICK"
	Wool   resourceCardType = "WOOL"
	Grain  resourceCardType = "GRAIN"
	Ore    resourceCardType = "ORE"
)
