package models

func NewHex(q int, r int) Hex {
	return Hex{
		q: q,
		r: r,
	}
}

func FindAdjacentHexesFromHexCorner(hexCorner HexCorner) []Hex {
	hexes := make([]Hex, 0)

	switch hexCorner.location {
	case Top:
		topLeftHex := NewHex(hexCorner.q, hexCorner.r-1)
		topRightHex := NewHex(hexCorner.q+1, hexCorner.r-1)
		bottomHex := NewHex(hexCorner.q, hexCorner.r)

		hexes = append(hexes, topLeftHex, topRightHex, bottomHex)
	case Bottom:
		topHex := NewHex(hexCorner.q, hexCorner.r)
		bottomLeftHex := NewHex(hexCorner.q-1, hexCorner.r+1)
		bottomRightHex := NewHex(hexCorner.q, hexCorner.r+1)

		hexes = append(hexes, topHex, bottomLeftHex, bottomRightHex)
	}

	return hexes
}

type Hex struct {
	q int
	r int
}

func (h Hex) GetQ() int {
	return h.q
}

func (h Hex) GetR() int {
	return h.r
}

func (h Hex) IsAdjacentWithHex(hex Hex) bool {
	hexDirections := []hexDirection{
		NewHexDirection(1, 0),
		NewHexDirection(0, 1),
		NewHexDirection(-1, 1),
		NewHexDirection(-1, 0),
		NewHexDirection(0, -1),
		NewHexDirection(1, -1),
	}

	for _, hexDirection := range hexDirections {
		if hexDirection.CalculateEndpoint(h) == hex {
			return true
		}
	}

	return false
}
