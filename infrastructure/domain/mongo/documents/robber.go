package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Robber struct {
	Document  `bson:",inline"`
	TerrainID primitive.ObjectID `bson:"terrain_id"`
	IsMoving  bool               `bson:"is_moving"`
}
