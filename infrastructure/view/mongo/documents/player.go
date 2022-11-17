package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID                 primitive.ObjectID `bson:"_id"`
	UserID             primitive.ObjectID `bson:"user_id"`
	Color              string             `bson:"color"`
	TurnOrder          int                `bson:"turn_order"`
	ReceivedOffer      bool               `bson:"received_offer"`
	DiscardedResources bool               `bson:"discarded_resources"`
	Score              int                `bson:"score"`
	Achievements       []*Achievement     `bson:"achievements"`
	ResourceCards      []*ResourceCard    `bson:"resource_cards"`
	DevelopmentCards   []*DevelopmentCard `bson:"development_cards"`
	Constructions      []*Construction    `bson:"constructions"`
	Roads              []*Road            `bson:"roads"`
}
