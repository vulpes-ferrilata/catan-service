package models

type ResourceCard struct {
	aggregate
	resourceCardType ResourceCardType
	offering         bool
}

func (r ResourceCard) GetType() ResourceCardType {
	return r.resourceCardType
}

func (r ResourceCard) IsOffering() bool {
	return r.offering
}
