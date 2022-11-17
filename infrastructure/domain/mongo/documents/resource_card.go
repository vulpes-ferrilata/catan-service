package documents

type ResourceCard struct {
	Document `bson:",inline"`
	Type     string `bson:"type"`
	Offering bool   `bson:"offering"`
}
