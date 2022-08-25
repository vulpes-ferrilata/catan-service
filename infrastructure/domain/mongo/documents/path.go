package documents

type Path struct {
	Document `bson:",inline"`
	Q        int    `bson:"q"`
	R        int    `bson:"r"`
	Location string `bson:"location"`
}
