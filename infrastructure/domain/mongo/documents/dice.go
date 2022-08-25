package documents

type Dice struct {
	Document `bson:",inline"`
	Number   int `bson:"number"`
}
