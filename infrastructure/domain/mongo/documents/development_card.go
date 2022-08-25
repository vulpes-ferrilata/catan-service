package documents

type DevelopmentCard struct {
	Document `bson:",inline"`
	Type     string `bson:"type"`
	Status   string `bson:"status"`
}
