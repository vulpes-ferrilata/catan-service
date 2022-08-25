package models

type ResourceCard struct {
	aggregate
	resourceCardType resourceCardType
	isSelected       bool
}

func (r ResourceCard) GetType() resourceCardType {
	return r.resourceCardType
}

func (r ResourceCard) IsSelected() bool {
	return r.isSelected
}
