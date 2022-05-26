package models

import (
	"github.com/VulpesFerrilata/catan-service/domain/models/common"
	"github.com/google/uuid"
)

func NewResourceCard(id uuid.UUID, resourceCardType ResourceCardType) *ResourceCard {
	return &ResourceCard{
		Entity:           common.NewEntity(id),
		resourceCardType: resourceCardType,
	}
}

type ResourceCard struct {
	common.Entity
	resourceCardType ResourceCardType
}

func (r ResourceCard) GetType() ResourceCardType {
	return r.resourceCardType
}
