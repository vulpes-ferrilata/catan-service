package models

import (
	"github.com/pkg/errors"
)

func NewAchievementType(value string) (achievementType, error) {
	achievementType := achievementType(value)

	switch achievementType {
	case LongestRoad, LargestArmy:
		return achievementType, nil
	default:
		return achievementType, errors.New("achievement type is invalid")
	}
}

type achievementType string

const (
	LongestRoad achievementType = "LONGEST_ROAD"
	LargestArmy achievementType = "LARGEST_ARMY"
)
