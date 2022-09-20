package documents

type Harbor struct {
	Document `bson:",inline"`
	Q        int    `bson:"q"`
	R        int    `bson:"r"`
	Type     string `bson:"type"`
}
