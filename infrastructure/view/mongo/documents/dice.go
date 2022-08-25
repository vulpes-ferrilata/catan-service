package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dice struct {
	ID       primitive.ObjectID `bson:"_id"`
	Number   int                `bson:"number"`
	IsRolled bool               `bson:"is_rolled"`
}
