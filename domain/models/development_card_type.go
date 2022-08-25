package models

import (
	"github.com/pkg/errors"
)

func NewDevelopmentCardType(value string) (developmentCardType, error) {
	developmentCardType := developmentCardType(value)

	switch developmentCardType {
	case Knight, Monopoly, RoadBuilding, YearOfPlenty, VictoryPoints:
		return developmentCardType, nil
	default:
		return developmentCardType, errors.New("development type is invalid")
	}
}

type developmentCardType string

const (
	Knight        developmentCardType = "KNIGHT"
	Monopoly      developmentCardType = "MONOPOLY"
	RoadBuilding  developmentCardType = "ROAD_BUILDING"
	YearOfPlenty  developmentCardType = "YEAR_OF_PLENTY"
	VictoryPoints developmentCardType = "VICTORY_POINTS"
)
