package models

import (
	"github.com/vulpes-ferrilata/slices"
)

func NewHexCorner(q int, r int, location HexCornerLocation) HexCorner {
	return HexCorner{
		q:        q,
		r:        r,
		location: location,
	}
}

func findAdjacentHexCornersFromHex(hex Hex) []HexCorner {
	hexCorners := make([]HexCorner, 0)

	topHexCorner := HexCorner{hex.q, hex.r, Top}
	topRightHexCorner := HexCorner{hex.q + 1, hex.r - 1, Bottom}
	bottomRightHexCorner := HexCorner{hex.q, hex.r + 1, Top}
	bottomHexCorner := HexCorner{hex.q, hex.r, Bottom}
	bottomLeftHexCorner := HexCorner{hex.q - 1, hex.r + 1, Top}
	topLeftHexCorner := HexCorner{hex.q, hex.r - 1, Bottom}

	hexCorners = append(
		hexCorners,
		topHexCorner,
		topRightHexCorner,
		bottomRightHexCorner,
		bottomHexCorner,
		bottomLeftHexCorner,
		topLeftHexCorner,
	)

	return hexCorners
}

func findAdjacentHexCornersFromHexEdge(hexEdge HexEdge) []HexCorner {
	hexCorners := make([]HexCorner, 0)

	var topHexCorner HexCorner
	var bottomHexCorner HexCorner

	switch hexEdge.location {
	case TopLeft:
		topHexCorner = HexCorner{hexEdge.q, hexEdge.r, Top}
		bottomHexCorner = HexCorner{hexEdge.q, hexEdge.r - 1, Bottom}
	case MiddleLeft:
		topHexCorner = HexCorner{hexEdge.q, hexEdge.r - 1, Bottom}
		bottomHexCorner = HexCorner{hexEdge.q - 1, hexEdge.r + 1, Top}
	case BottomLeft:
		topHexCorner = HexCorner{hexEdge.q - 1, hexEdge.r + 1, Top}
		bottomHexCorner = HexCorner{hexEdge.q, hexEdge.r, Bottom}
	}

	hexCorners = append(
		hexCorners,
		topHexCorner,
		bottomHexCorner,
	)

	return hexCorners
}

func findAdjacentHexCornersFromHexCorner(hexCorner HexCorner) []HexCorner {
	hexCorners := make([]HexCorner, 0)

	switch hexCorner.location {
	case Top:
		topHexCorner := HexCorner{hexCorner.q + 1, hexCorner.r - 2, Bottom}
		bottomRightHexCorner := HexCorner{hexCorner.q + 1, hexCorner.r - 1, Bottom}
		bottomLeftHexCorner := HexCorner{hexCorner.q, hexCorner.r - 1, Bottom}

		hexCorners = append(hexCorners, topHexCorner, bottomRightHexCorner, bottomLeftHexCorner)
	case Bottom:
		topRightHexCorner := HexCorner{hexCorner.q, hexCorner.r + 1, Top}
		bottomHexCorner := HexCorner{hexCorner.q - 1, hexCorner.r + 2, Top}
		topLeftHexCorner := HexCorner{hexCorner.q - 1, hexCorner.r + 1, Top}

		hexCorners = append(hexCorners, topRightHexCorner, bottomHexCorner, topLeftHexCorner)
	}

	return hexCorners
}

func findIntersectionHexCornerBetweenTwoHexEdges(hexEdge1 HexEdge, hexEdge2 HexEdge) (HexCorner, error) {
	adjacentHexCorners := findAdjacentHexCornersFromHexEdge(hexEdge1)

	return slices.Find(func(adjacentHexCorner HexCorner) (bool, error) {
		return adjacentHexCorner.isAdjacentWithHexEdge(hexEdge2)
	}, adjacentHexCorners...)
}

type HexCorner struct {
	q        int
	r        int
	location HexCornerLocation
}

func (h HexCorner) GetQ() int {
	return h.q
}

func (h HexCorner) GetR() int {
	return h.r
}

func (h HexCorner) GetLocation() HexCornerLocation {
	return h.location
}

func (h HexCorner) isAdjacentWithHex(hex Hex) (bool, error) {
	adjacentHexes := findAdjacentHexesFromHexCorner(h)

	return slices.Any(func(adjacentHex Hex) (bool, error) {
		return adjacentHex == hex, nil
	}, adjacentHexes...)
}

func (h HexCorner) isAdjacentWithHexEdge(hexEdge HexEdge) (bool, error) {
	adjacentHexEdges := findAdjacentHexEdgesFromHexCorner(h)

	return slices.Any(func(adjacentHexEdge HexEdge) (bool, error) {
		return adjacentHexEdge == hexEdge, nil
	}, adjacentHexEdges...)
}

func (h HexCorner) isAdjacentWithHexCorner(hexCorner HexCorner) (bool, error) {
	adjacentHexCorners := findAdjacentHexCornersFromHexCorner(h)

	return slices.Any(func(adjacentHexCorner HexCorner) (bool, error) {
		return adjacentHexCorner == hexCorner, nil
	}, adjacentHexCorners...)
}
