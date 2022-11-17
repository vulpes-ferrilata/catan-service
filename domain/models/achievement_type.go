package models

import (
	"github.com/pkg/errors"
)

func NewAchievementType(enum string) (AchievementType, error) {
	achievementType := AchievementType{enum}

	switch achievementType {
	case LongestRoad, LargestArmy:
		return achievementType, nil
	default:
		return achievementType, errors.New("achievement type is invalid")
	}
}

type AchievementType struct {
	enum string
}

func (a AchievementType) String() string {
	return a.enum
}

var (
	LongestRoad = AchievementType{"LongestRoad"}
	LargestArmy = AchievementType{"LargestArmy"}
)
