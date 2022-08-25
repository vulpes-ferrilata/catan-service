package documents

type Road struct {
	Document `bson:",inline"`
	Path     *Path `bson:"path"`
}
