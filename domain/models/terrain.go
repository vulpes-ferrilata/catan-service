package models

type Terrain struct {
	aggregate
	hex         Hex
	number      int
	terrainType terrainType
}

func (t Terrain) GetHex() Hex {
	return t.hex
}

func (t Terrain) GetNumber() int {
	return t.number
}

func (t Terrain) GetType() terrainType {
	return t.terrainType
}
