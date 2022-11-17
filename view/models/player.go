package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID                 primitive.ObjectID
	UserID             primitive.ObjectID
	Color              string
	TurnOrder          int
	ReceivedOffer      bool
	DiscardedResources bool
	Score              int
	Achievements       []*Achievement
	ResourceCards      []*ResourceCard
	DevelopmentCards   []*DevelopmentCard
	Constructions      []*Construction
	Roads              []*Road
}
