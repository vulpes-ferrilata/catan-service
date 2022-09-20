package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	ID               primitive.ObjectID `bson:"_id"`
	Status           string             `bson:"status"`
	Phase            string             `bson:"phase"`
	Turn             int                `bson:"turn"`
	Players          []*Player          `bson:"players"`
	Dices            []*Dice            `bson:"dices"`
	Achievements     []*Achievement     `bson:"achievements"`
	ResourceCards    []*ResourceCard    `bson:"resource_cards"`
	DevelopmentCards []*DevelopmentCard `bson:"development_cards"`
	Terrains         []*Terrain         `bson:"terrains"`
	Lands            []*Land            `bson:"lands"`
	Paths            []*Path            `bson:"paths"`
}
