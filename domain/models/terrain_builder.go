package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TerrainBuilder struct {
	id          primitive.ObjectID
	hex         Hex
	number      int
	terrainType terrainType
	harbor      *Harbor
	robber      *Robber
}

func (t TerrainBuilder) SetID(id primitive.ObjectID) TerrainBuilder {
	t.id = id

	return t
}
func (t TerrainBuilder) SetHex(hex Hex) TerrainBuilder {
	t.hex = hex

	return t
}

func (t TerrainBuilder) SetNumber(number int) TerrainBuilder {
	t.number = number

	return t
}

func (t TerrainBuilder) SetType(terrainType terrainType) TerrainBuilder {
	t.terrainType = terrainType

	return t
}

func (t TerrainBuilder) SetHarbor(harbor *Harbor) TerrainBuilder {
	t.harbor = harbor

	return t
}

func (t TerrainBuilder) SetRobber(robber *Robber) TerrainBuilder {
	t.robber = robber

	return t
}

func (t TerrainBuilder) Create() *Terrain {
	return &Terrain{
		aggregate: aggregate{
			id: t.id,
		},
		hex:         t.hex,
		number:      t.number,
		terrainType: t.terrainType,
		harbor:      t.harbor,
		robber:      t.robber,
	}
}
