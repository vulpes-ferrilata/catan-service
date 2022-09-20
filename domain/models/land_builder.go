package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LandBuilder struct {
	id        primitive.ObjectID
	hexCorner HexCorner
}

func (l LandBuilder) SetID(id primitive.ObjectID) LandBuilder {
	l.id = id

	return l
}

func (l LandBuilder) SetHexCorner(hexCorner HexCorner) LandBuilder {
	l.hexCorner = hexCorner

	return l
}
func (l LandBuilder) Create() *Land {
	return &Land{
		aggregate: aggregate{
			id: l.id,
		},
		hexCorner: l.hexCorner,
	}
}
