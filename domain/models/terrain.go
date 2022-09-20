package models

type Terrain struct {
	aggregate
	hex         Hex
	number      int
	terrainType terrainType
	harbor      *Harbor
	robber      *Robber
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

func (t Terrain) GetHarbor() *Harbor {
	return t.harbor
}

func (t Terrain) GetRobber() *Robber {
	return t.robber
}
