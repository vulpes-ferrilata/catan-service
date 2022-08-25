package models

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	aggregate
	userID           primitive.ObjectID
	color            playerColor
	turnOrder        int
	isActive         bool
	isOffered        bool
	achievements     []*Achievement
	resourceCards    []*ResourceCard
	developmentCards []*DevelopmentCard
	constructions    []*Construction
	roads            []*Road
}

func (p Player) GetUserID() primitive.ObjectID {
	return p.userID
}

func (p Player) GetColor() playerColor {
	return p.color
}

func (p Player) GetTurnOrder() int {
	return p.turnOrder
}

func (p Player) IsActive() bool {
	return p.isActive
}

func (p Player) IsOffered() bool {
	return p.isOffered
}

func (p Player) GetAchievements() []*Achievement {
	return slices.Clone(p.achievements)
}

func (p Player) GetResourceCards() []*ResourceCard {
	return slices.Clone(p.resourceCards)
}

func (p Player) GetDevelopmentCards() []*DevelopmentCard {
	return slices.Clone(p.developmentCards)
}

func (p Player) GetConstructions() []*Construction {
	return slices.Clone(p.constructions)
}

func (p Player) GetRoads() []*Road {
	return slices.Clone(p.roads)
}
