package models

type ResourceCard struct {
	aggregate
	resourceCardType ResourceCardType
	isSelected       bool
}

func (r ResourceCard) GetType() ResourceCardType {
	return r.resourceCardType
}

func (r ResourceCard) IsSelected() bool {
	return r.isSelected
}
