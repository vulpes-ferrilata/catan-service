package models

import "github.com/pkg/errors"

func NewResourceCardType(value string) (ResourceCardType, error) {
	resourceCardType := ResourceCardType(value)

	switch resourceCardType {
	case Lumber, Brick, Wool, Grain, Ore:
		return resourceCardType, nil
	default:
		return resourceCardType, errors.Wrap(ErrEnumIsInvalid, value)
	}
}

type ResourceCardType string

const (
	Lumber ResourceCardType = "Lumber"
	Brick  ResourceCardType = "Brick"
	Wool   ResourceCardType = "Wool"
	Grain  ResourceCardType = "Grain"
	Ore    ResourceCardType = "Ore"
)
