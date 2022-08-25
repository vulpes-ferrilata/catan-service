package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Harbor struct {
	aggregate
	terrainID  primitive.ObjectID
	hex        Hex
	harborType harborType
}

func (t Harbor) GetTerrainID() primitive.ObjectID {
	return t.terrainID
}

func (t Harbor) GetHex() Hex {
	return t.hex
}

func (t Harbor) GetType() harborType {
	return t.harborType
}
