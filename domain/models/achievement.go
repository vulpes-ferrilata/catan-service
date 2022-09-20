package models

type Achievement struct {
	aggregate
	achievementType AchievementType
}

func (a Achievement) GetType() AchievementType {
	return a.achievementType
}
