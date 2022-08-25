package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LandBuilder interface {
	SetID(id primitive.ObjectID) LandBuilder
	SetHexCorner(hexCorner HexCorner) LandBuilder
	Create() *Land
}

func NewLandBuilder() LandBuilder {
	return &landBuilder{}
}

type landBuilder struct {
	id        primitive.ObjectID
	hexCorner HexCorner
}

func (l *landBuilder) SetID(id primitive.ObjectID) LandBuilder {
	l.id = id

	return l
}

func (l *landBuilder) SetHexCorner(hexCorner HexCorner) LandBuilder {
	l.hexCorner = hexCorner

	return l
}
func (l landBuilder) Create() *Land {
	return &Land{
		aggregate: aggregate{
			id: l.id,
		},
		hexCorner: l.hexCorner,
	}
}
