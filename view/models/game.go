package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	ID               primitive.ObjectID
	Status           string
	Turn             int
	IsRolledDices    bool
	Players          []*Player
	Dices            []*Dice
	Achievements     []*Achievement
	ResourceCards    []*ResourceCard
	DevelopmentCards []*DevelopmentCard
	Terrains         []*Terrain
	Harbors          []*Harbor
	Robber           *Robber
	Lands            []*Land
	Paths            []*Path
	Version          int
}
