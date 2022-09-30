package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	ID               primitive.ObjectID
	Status           string
	Phase            string
	Turn             int
	ActivePlayer     *Player
	Players          []*Player
	Dices            []*Dice
	Achievements     []*Achievement
	ResourceCards    []*ResourceCard
	DevelopmentCards []*DevelopmentCard
	Terrains         []*Terrain
	Lands            []*Land
	Paths            []*Path
	Version          int
}
