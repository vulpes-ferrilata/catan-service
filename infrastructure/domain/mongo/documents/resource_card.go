package documents

type ResourceCard struct {
	Document   `bson:",inline"`
	Type       string `bson:"type"`
	IsSelected bool   `bson:"is_selected"`
}
