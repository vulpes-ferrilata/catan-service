package models

import "github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"

func NewHex(q int, r int) Hex {
	return Hex{
		q: q,
		r: r,
	}
}

func findAdjacentHexesFromHex(hex Hex) []Hex {
	topRightHex := Hex{hex.q + 1, hex.r - 1}
	rightHex := Hex{hex.q + 1, hex.r}
	bottomRightHex := Hex{hex.q, hex.r + 1}
	bottomLeftHex := Hex{hex.q - 1, hex.r + 1}
	leftHex := Hex{hex.q - 1, hex.r}
	topLeftHex := Hex{hex.q, hex.r - 1}

	return []Hex{
		topRightHex,
		rightHex,
		bottomRightHex,
		bottomLeftHex,
		leftHex,
		topLeftHex,
	}
}

func findAdjacentHexesFromHexCorner(hexCorner HexCorner) []Hex {
	hexes := make([]Hex, 0)

	switch hexCorner.location {
	case Top:
		topLeftHex := Hex{hexCorner.q, hexCorner.r - 1}
		topRightHex := Hex{hexCorner.q + 1, hexCorner.r - 1}
		bottomHex := Hex{hexCorner.q, hexCorner.r}

		hexes = append(hexes, topLeftHex, topRightHex, bottomHex)
	case Bottom:
		topHex := Hex{hexCorner.q, hexCorner.r}
		bottomLeftHex := Hex{hexCorner.q - 1, hexCorner.r + 1}
		bottomRightHex := Hex{hexCorner.q, hexCorner.r + 1}

		hexes = append(hexes, topHex, bottomLeftHex, bottomRightHex)
	}

	return hexes
}

func findIntersectionHexCornersBetweenTwoHexes(hex1 Hex, hex2 Hex) []HexCorner {
	adjacentHexCorners := findAdjacentHexCornersFromHex(hex1)

	return slices.Filter(func(adjacentHexCorner HexCorner) bool {
		return adjacentHexCorner.isAdjacentWithHex(hex2)
	}, adjacentHexCorners)
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

func (h Hex) isAdjacentWithHex(hex Hex) bool {
	adjacentHexes := findAdjacentHexesFromHex(h)

	return slices.Contains(adjacentHexes, hex)
}
