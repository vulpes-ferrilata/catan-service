package documents

type Construction struct {
	Document `bson:",inline"`
	Type     string `bson:"type"`
	Land     *Land  `bson:"land"`
}
