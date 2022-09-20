package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PathBuilder struct {
	id      primitive.ObjectID
	hexEdge HexEdge
}

func (p PathBuilder) SetID(id primitive.ObjectID) PathBuilder {
	p.id = id

	return p
}

func (p PathBuilder) SetHexEdge(hexEdge HexEdge) PathBuilder {
	p.hexEdge = hexEdge

	return p
}

func (p PathBuilder) Create() *Path {
	return &Path{
		aggregate: aggregate{
			id: p.id,
		},
		hexEdge: p.hexEdge,
	}
}
