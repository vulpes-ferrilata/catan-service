package models

type Achievement struct {
	aggregate
	achievementType achievementType
}

func (a Achievement) GetType() achievementType {
	return a.achievementType
}
