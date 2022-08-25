package models

type Path struct {
	aggregate
	hexEdge HexEdge
}

func (r Path) GetHexEdge() HexEdge {
	return r.hexEdge
}
