package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HarborBuilder interface {
	SetID(id primitive.ObjectID) HarborBuilder
	SetTerrainID(terrainID primitive.ObjectID) HarborBuilder
	SetHex(hex Hex) HarborBuilder
	SetHarborType(harborType harborType) HarborBuilder
	Create() *Harbor
}

func NewHarborBuilder() HarborBuilder {
	return &harborBuilder{}
}

type harborBuilder struct {
	id         primitive.ObjectID
	terrainID  primitive.ObjectID
	hex        Hex
	harborType harborType
}

func (h *harborBuilder) SetID(id primitive.ObjectID) HarborBuilder {
	h.id = id

	return h
}

func (h *harborBuilder) SetTerrainID(terrainID primitive.ObjectID) HarborBuilder {
	h.terrainID = terrainID

	return h
}

func (h *harborBuilder) SetHex(hex Hex) HarborBuilder {
	h.hex = hex

	return h
}

func (h *harborBuilder) SetHarborType(harborType harborType) HarborBuilder {
	h.harborType = harborType

	return h
}

func (h harborBuilder) Create() *Harbor {
	return &Harbor{
		aggregate: aggregate{
			id: h.id,
		},
		terrainID:  h.terrainID,
		hex:        h.hex,
		harborType: h.harborType,
	}
}
