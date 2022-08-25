package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TerrainBuilder interface {
	SetID(id primitive.ObjectID) TerrainBuilder
	SetHex(hex Hex) TerrainBuilder
	SetNumber(number int) TerrainBuilder
	SetType(terrainType terrainType) TerrainBuilder
	Create() *Terrain
}

func NewTerrainBuilder() TerrainBuilder {
	return &terrainBuilder{}
}

type terrainBuilder struct {
	id          primitive.ObjectID
	hex         Hex
	number      int
	terrainType terrainType
}

func (t *terrainBuilder) SetID(id primitive.ObjectID) TerrainBuilder {
	t.id = id

	return t
}
func (t *terrainBuilder) SetHex(hex Hex) TerrainBuilder {
	t.hex = hex

	return t
}

func (t *terrainBuilder) SetNumber(number int) TerrainBuilder {
	t.number = number

	return t
}

func (t *terrainBuilder) SetType(terrainType terrainType) TerrainBuilder {
	t.terrainType = terrainType

	return t
}

func (t terrainBuilder) Create() *Terrain {
	return &Terrain{
		aggregate: aggregate{
			id: t.id,
		},
		hex:         t.hex,
		number:      t.number,
		terrainType: t.terrainType,
	}
}
