package documents

type Game struct {
	DocumentRoot     `bson:",inline"`
	Status           string             `bson:"status"`
	Phase            string             `bson:"phase"`
	Turn             int                `bson:"turn"`
	ActivePlayer     *Player            `bson:"active_player"`
	Players          []*Player          `bson:"players"`
	Dices            []*Dice            `bson:"dices"`
	Achievements     []*Achievement     `bson:"achievements"`
	ResourceCards    []*ResourceCard    `bson:"resource_cards"`
	DevelopmentCards []*DevelopmentCard `bson:"development_cards"`
	Terrains         []*Terrain         `bson:"terrains"`
	Lands            []*Land            `bson:"lands"`
	Paths            []*Path            `bson:"paths"`
}
