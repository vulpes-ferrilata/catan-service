package models

import (
	"github.com/pkg/errors"
)

func NewDevelopmentCardType(enum string) (DevelopmentCardType, error) {
	developmentCardType := DevelopmentCardType{enum}

	switch developmentCardType {
	case Knight, Monopoly, RoadBuilding, YearOfPlenty, Chapel, GreatHall, Library, Market, University:
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
	Knight       = DevelopmentCardType{"Knight"}
	Monopoly     = DevelopmentCardType{"Monopoly"}
	RoadBuilding = DevelopmentCardType{"RoadBuilding"}
	YearOfPlenty = DevelopmentCardType{"YearOfPlenty"}
	Chapel       = DevelopmentCardType{"Chapel"}
	GreatHall    = DevelopmentCardType{"GreatHall"}
	Library      = DevelopmentCardType{"Library"}
	Market       = DevelopmentCardType{"Market"}
	University   = DevelopmentCardType{"University"}
)
