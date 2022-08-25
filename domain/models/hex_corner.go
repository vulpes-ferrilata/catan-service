package models

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
)

func NewHexCorner(q int, r int, location hexCornerLocation) HexCorner {
	return HexCorner{
		q:        q,
		r:        r,
		location: location,
	}
}

func FindAdjacentHexCorners(hex Hex) []HexCorner {
	hexCorners := make([]HexCorner, 0)

	topHexCorner := NewHexCorner(hex.q, hex.r, Top)
	topRightHexCorner := NewHexCorner(hex.q+1, hex.r-1, Bottom)
	bottomRightHexCorner := NewHexCorner(hex.q, hex.r+1, Top)
	bottomHexCorner := NewHexCorner(hex.q, hex.r, Bottom)
	bottomLeftHexCorner := NewHexCorner(hex.q-1, hex.r+1, Top)
	topLeftHexCorner := NewHexCorner(hex.q, hex.r-1, Bottom)

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

func FindAdjacentHexCornersFromHexEdge(hexEdge HexEdge) []HexCorner {
	hexCorners := make([]HexCorner, 0)

	var topHexCorner HexCorner
	var bottomHexCorner HexCorner

	switch hexEdge.location {
	case TopLeft:
		topHexCorner = NewHexCorner(hexEdge.q, hexEdge.r, Top)
		bottomHexCorner = NewHexCorner(hexEdge.q, hexEdge.r-1, Bottom)
	case MiddleLeft:
		topHexCorner = NewHexCorner(hexEdge.q, hexEdge.r-1, Bottom)
		bottomHexCorner = NewHexCorner(hexEdge.q-1, hexEdge.r+1, Top)
	case BottomLeft:
		topHexCorner = NewHexCorner(hexEdge.q-1, hexEdge.r+1, Top)
		bottomHexCorner = NewHexCorner(hexEdge.q, hexEdge.r, Bottom)
	}

	hexCorners = append(
		hexCorners,
		topHexCorner,
		bottomHexCorner,
	)

	return hexCorners
}

func FindAdjacentHexCornersFromHexCorner(hexCorner HexCorner) []HexCorner {
	hexCorners := make([]HexCorner, 0)

	switch hexCorner.location {
	case Top:
		topHexCorner := NewHexCorner(hexCorner.q+1, hexCorner.r-2, Bottom)
		bottomRightHexCorner := NewHexCorner(hexCorner.q+1, hexCorner.r-1, Bottom)
		bottomLeftHexCorner := NewHexCorner(hexCorner.q, hexCorner.r-1, Bottom)

		hexCorners = append(hexCorners, topHexCorner, bottomRightHexCorner, bottomLeftHexCorner)
	case Bottom:
		topRightHexCorner := NewHexCorner(hexCorner.q, hexCorner.r+1, Top)
		bottomHexCorner := NewHexCorner(hexCorner.q-1, hexCorner.r+2, Top)
		topLeftHexCorner := NewHexCorner(hexCorner.q-1, hexCorner.r+1, Top)

		hexCorners = append(hexCorners, topRightHexCorner, bottomHexCorner, topLeftHexCorner)
	}

	return hexCorners
}

type HexCorner struct {
	q        int
	r        int
	location hexCornerLocation
}

func (h HexCorner) GetQ() int {
	return h.q
}

func (h HexCorner) GetR() int {
	return h.r
}

func (h HexCorner) GetLocation() hexCornerLocation {
	return h.location
}

func (h HexCorner) IsAdjacentWithHex(hex Hex) bool {
	adjacentHexes := FindAdjacentHexesFromHexCorner(h)

	return slices.Contains(adjacentHexes, hex)
}

func (h HexCorner) IsAdjacentWithHexEdge(hexEdge HexEdge) bool {
	adjacentHexEdges := FindAdjacentHexEdgesFromHexCorner(h)

	return slices.Contains(adjacentHexEdges, hexEdge)
}

func (h HexCorner) IsAdjacentWithHexCorner(hexCorner HexCorner) bool {
	adjacentHexCorners := FindAdjacentHexCornersFromHexCorner(h)

	return slices.Contains(adjacentHexCorners, hexCorner)
}
