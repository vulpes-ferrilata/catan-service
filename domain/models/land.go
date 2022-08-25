package models

type Land struct {
	aggregate
	hexCorner HexCorner
}

func (l Land) GetHexCorner() HexCorner {
	return l.hexCorner
}
