package documents

type Pagination[Document any] struct {
	Metadata `bson:"metadata"`
	Data     []Document `bson:"data"`
}
