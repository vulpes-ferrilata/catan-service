package models

import (
	"github.com/pkg/errors"
)

func NewDevelopmentCardType(enum string) (DevelopmentCardType, error) {
	developmentCardType := DevelopmentCardType{enum}

	switch developmentCardType {
	case Knight, Monopoly, RoadBuilding, YearOfPlenty, VictoryPoints:
		return developmentCardType, nil
	default:
		return developmentCardType, errors.New("development type is invalid")
	}
}

type DevelopmentCardType struct {
	enum string
}

func (d DevelopmentCardType) String() string {
	return d.enum
}

var (
	Knight        = DevelopmentCardType{"KNIGHT"}
	Monopoly      = DevelopmentCardType{"MONOPOLY"}
	RoadBuilding  = DevelopmentCardType{"ROAD_BUILDING"}
	YearOfPlenty  = DevelopmentCardType{"YEAR_OF_PLENTY"}
	VictoryPoints = DevelopmentCardType{"VICTORY_POINTS"}
)
