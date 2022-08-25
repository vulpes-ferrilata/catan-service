package documents

type Game struct {
	DocumentRoot     `bson:",inline"`
	Status           string             `bson:"status"`
	Turn             int                `bson:"turn"`
	IsRolledDices    bool               `bson:"is_rolled_dices"`
	Players          []*Player          `bson:"players"`
	Dices            []*Dice            `bson:"dices"`
	Achievements     []*Achievement     `bson:"achievements"`
	ResourceCards    []*ResourceCard    `bson:"resource_cards"`
	DevelopmentCards []*DevelopmentCard `bson:"development_cards"`
	Terrains         []*Terrain         `bson:"terrains"`
	Harbors          []*Harbor          `bson:"harbors"`
	Robber           *Robber            `bson:"robber"`
	Lands            []*Land            `bson:"lands"`
	Paths            []*Path            `bson:"paths"`
}
