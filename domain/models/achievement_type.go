package models

import "github.com/pkg/errors"

func NewAchievementType(value string) (AchievementType, error) {
	achievementType := AchievementType(value)

	switch achievementType {
	case LongestRoad, LargestArmy:
		return achievementType, nil
	default:
		return achievementType, errors.Wrap(ErrEnumIsInvalid, value)
	}
}

type AchievementType string

const (
	LongestRoad AchievementType = "LongestRoad"
	LargestArmy AchievementType = "LargestArmy"
)
