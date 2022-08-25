package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PathBuilder interface {
	SetID(id primitive.ObjectID) PathBuilder
	SetHexEdge(hexEdge HexEdge) PathBuilder
	Create() *Path
}

func NewPathBuilder() PathBuilder {
	return &pathBuilder{}
}

type pathBuilder struct {
	id      primitive.ObjectID
	hexEdge HexEdge
}

func (p *pathBuilder) SetID(id primitive.ObjectID) PathBuilder {
	p.id = id

	return p
}

func (p *pathBuilder) SetHexEdge(hexEdge HexEdge) PathBuilder {
	p.hexEdge = hexEdge

	return p
}

func (p pathBuilder) Create() *Path {
	return &Path{
		aggregate: aggregate{
			id: p.id,
		},
		hexEdge: p.hexEdge,
	}
}
