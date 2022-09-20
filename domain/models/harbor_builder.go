package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HarborBuilder struct {
	id         primitive.ObjectID
	hex        Hex
	harborType HarborType
}

func (h HarborBuilder) SetID(id primitive.ObjectID) HarborBuilder {
	h.id = id

	return h
}

func (h HarborBuilder) SetHex(hex Hex) HarborBuilder {
	h.hex = hex

	return h
}

func (h HarborBuilder) SetType(harborType HarborType) HarborBuilder {
	h.harborType = harborType

	return h
}

func (h HarborBuilder) Create() *Harbor {
	return &Harbor{
		aggregate: aggregate{
			id: h.id,
		},
		hex:        h.hex,
		harborType: h.harborType,
	}
}
