package documents

type Terrain struct {
	Document `bson:",inline"`
	Q        int     `bson:"q"`
	R        int     `bson:"r"`
	Number   int     `bson:"number"`
	Type     string  `bson:"type"`
	Harbor   *Harbor `bson:"harbor"`
	Robber   *Robber `bson:"robber"`
}
