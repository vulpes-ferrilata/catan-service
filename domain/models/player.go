package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	aggregate
	userID             primitive.ObjectID
	color              PlayerColor
	turnOrder          int
	receivedOffer      bool
	discardedResources bool
	score              int
	achievements       []*Achievement
	resourceCards      []*ResourceCard
	developmentCards   []*DevelopmentCard
	constructions      []*Construction
	roads              []*Road
}

func (p Player) GetUserID() primitive.ObjectID {
	return p.userID
}

func (p Player) GetColor() PlayerColor {
	return p.color
}

func (p Player) GetTurnOrder() int {
	return p.turnOrder
}

func (p Player) IsReceivedOffer() bool {
	return p.receivedOffer
}

func (p Player) IsDiscardedResources() bool {
	return p.discardedResources
}

func (p Player) GetScore() int {
	return p.score
}

func (p Player) GetAchievements() []*Achievement {
	return p.achievements
}

func (p Player) GetResourceCards() []*ResourceCard {
	return p.resourceCards
}

func (p Player) GetDevelopmentCards() []*DevelopmentCard {
	return p.developmentCards
}

func (p Player) GetConstructions() []*Construction {
	return p.constructions
}

func (p Player) GetRoads() []*Road {
	return p.roads
}
