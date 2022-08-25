package models

import "github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"

func NewHexEdge(q int, r int, location hexEdgeLocation) HexEdge {
	return HexEdge{
		q:        q,
		r:        r,
		location: location,
	}
}

func FindAdjacentHexEdges(hex Hex) []HexEdge {
	hexEdges := make([]HexEdge, 0)

	topRightHexEdge := NewHexEdge(hex.q+1, hex.r-1, BottomLeft)
	middleRightHexEdge := NewHexEdge(hex.q+1, hex.r, MiddleLeft)
	bottomRightHexEdge := NewHexEdge(hex.q, hex.r+1, TopLeft)
	bottomLeftHexEdge := NewHexEdge(hex.q, hex.r, BottomLeft)
	middleLeftHexEdge := NewHexEdge(hex.q, hex.r, MiddleLeft)
	topLeftHexEdge := NewHexEdge(hex.q, hex.r, TopLeft)

	hexEdges = append(
		hexEdges,
		topRightHexEdge,
		middleRightHexEdge,
		bottomRightHexEdge,
		bottomLeftHexEdge,
		middleLeftHexEdge,
		topLeftHexEdge,
	)

	return hexEdges
}

func FindAdjacentHexEdgesFromHexCorner(hexCorner HexCorner) []HexEdge {
	hexEdges := make([]HexEdge, 0)

	switch hexCorner.location {
	case Top:
		topHexEdge := NewHexEdge(hexCorner.q+1, hexCorner.r-1, MiddleLeft)
		bottomRightHexEdge := NewHexEdge(hexCorner.q+1, hexCorner.r-1, BottomLeft)
		bottomLeftHexEdge := NewHexEdge(hexCorner.q, hexCorner.r, TopLeft)

		hexEdges = append(hexEdges, topHexEdge, bottomRightHexEdge, bottomLeftHexEdge)
	case Bottom:
		topRightHexEdge := NewHexEdge(hexCorner.q+1, hexCorner.r+1, TopLeft)
		bottomHexEdge := NewHexEdge(hexCorner.q+1, hexCorner.r+1, MiddleLeft)
		topLeftHexEdge := NewHexEdge(hexCorner.q, hexCorner.r, BottomLeft)

		hexEdges = append(hexEdges, topRightHexEdge, bottomHexEdge, topLeftHexEdge)
	}

	return hexEdges
}

func FindAdjacentHexEdgesFromHexEdge(hexEdge HexEdge) []HexEdge {
	hexEdges := make([]HexEdge, 0)

	var topLeftHexEdge HexEdge
	var topRightHexEdge HexEdge
	var bottomLeftHexEdge HexEdge
	var bottomRightHexEdge HexEdge

	switch hexEdge.location {
	case TopLeft:
		topLeftHexEdge = NewHexEdge(hexEdge.q+1, hexEdge.r-1, MiddleLeft)
		topRightHexEdge = NewHexEdge(hexEdge.q+1, hexEdge.r-1, BottomLeft)
		bottomLeftHexEdge = NewHexEdge(hexEdge.q, hexEdge.r-1, BottomLeft)
		bottomRightHexEdge = NewHexEdge(hexEdge.q, hexEdge.r, MiddleLeft)
	case MiddleLeft:
		topLeftHexEdge = NewHexEdge(hexEdge.q, hexEdge.r-1, BottomLeft)
		topRightHexEdge = NewHexEdge(hexEdge.q, hexEdge.r, TopLeft)
		bottomLeftHexEdge = NewHexEdge(hexEdge.q-1, hexEdge.r+1, TopLeft)
		bottomRightHexEdge = NewHexEdge(hexEdge.q, hexEdge.r, BottomLeft)
	case BottomLeft:
		topLeftHexEdge = NewHexEdge(hexEdge.q-1, hexEdge.r+1, TopLeft)
		topRightHexEdge = NewHexEdge(hexEdge.q, hexEdge.r, MiddleLeft)
		bottomLeftHexEdge = NewHexEdge(hexEdge.q, hexEdge.r+1, MiddleLeft)
		bottomRightHexEdge = NewHexEdge(hexEdge.q, hexEdge.r+1, TopLeft)
	}

	hexEdges = append(
		hexEdges,
		topLeftHexEdge,
		topRightHexEdge,
		bottomLeftHexEdge,
		bottomRightHexEdge,
	)

	return hexEdges
}

type HexEdge struct {
	q        int
	r        int
	location hexEdgeLocation
}

func (h HexEdge) GetQ() int {
	return h.q
}

func (h HexEdge) GetR() int {
	return h.r
}

func (h HexEdge) GetLocation() hexEdgeLocation {
	return h.location
}

func (h HexEdge) IsAdjacentWithHexCorner(hexCorner HexCorner) bool {
	adjacentHexCorners := FindAdjacentHexCornersFromHexEdge(h)

	return slices.Contains(adjacentHexCorners, hexCorner)
}

func (h HexEdge) IsAdjacentWithHexEdge(hexEdge HexEdge) bool {
	adjacentHexEdges := FindAdjacentHexEdgesFromHexEdge(h)

	return slices.Contains(adjacentHexEdges, hexEdge)
}
