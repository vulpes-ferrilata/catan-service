package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Harbor struct {
	Document  `bson:",inline"`
	TerrainID primitive.ObjectID `bson:"terrain_id"`
	Q         int                `bson:"q"`
	R         int                `bson:"r"`
	Type      string             `bson:"type"`
}
