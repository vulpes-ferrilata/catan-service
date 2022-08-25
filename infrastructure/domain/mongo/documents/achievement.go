package documents

type Achievement struct {
	Document `bson:",inline"`
	Type     string `bson:"type"`
}
