package models

import "github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"

func NewHexEdge(q int, r int, location HexEdgeLocation) HexEdge {
	return HexEdge{
		q:        q,
		r:        r,
		location: location,
	}
}

func findAdjacentHexEdgesFromHex(hex Hex) []HexEdge {
	hexEdges := make([]HexEdge, 0)

	topRightHexEdge := HexEdge{hex.q + 1, hex.r - 1, BottomLeft}
	middleRightHexEdge := HexEdge{hex.q + 1, hex.r, MiddleLeft}
	bottomRightHexEdge := HexEdge{hex.q, hex.r + 1, TopLeft}
	bottomLeftHexEdge := HexEdge{hex.q, hex.r, BottomLeft}
	middleLeftHexEdge := HexEdge{hex.q, hex.r, MiddleLeft}
	topLeftHexEdge := HexEdge{hex.q, hex.r, TopLeft}

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

func findAdjacentHexEdgesFromHexCorner(hexCorner HexCorner) []HexEdge {
	hexEdges := make([]HexEdge, 0)

	switch hexCorner.location {
	case Top:
		topHexEdge := HexEdge{hexCorner.q + 1, hexCorner.r - 1, MiddleLeft}
		bottomRightHexEdge := HexEdge{hexCorner.q + 1, hexCorner.r - 1, BottomLeft}
		bottomLeftHexEdge := HexEdge{hexCorner.q, hexCorner.r, TopLeft}

		hexEdges = append(hexEdges, topHexEdge, bottomRightHexEdge, bottomLeftHexEdge)
	case Bottom:
		topRightHexEdge := HexEdge{hexCorner.q, hexCorner.r + 1, TopLeft}
		bottomHexEdge := HexEdge{hexCorner.q, hexCorner.r + 1, MiddleLeft}
		topLeftHexEdge := HexEdge{hexCorner.q, hexCorner.r, BottomLeft}

		hexEdges = append(hexEdges, topRightHexEdge, bottomHexEdge, topLeftHexEdge)
	}

	return hexEdges
}

func findAdjacentHexEdgesFromHexEdge(hexEdge HexEdge) []HexEdge {
	hexEdges := make([]HexEdge, 0)

	var topLeftHexEdge HexEdge
	var topRightHexEdge HexEdge
	var bottomLeftHexEdge HexEdge
	var bottomRightHexEdge HexEdge

	switch hexEdge.location {
	case TopLeft:
		topLeftHexEdge = HexEdge{hexEdge.q + 1, hexEdge.r - 1, MiddleLeft}
		topRightHexEdge = HexEdge{hexEdge.q + 1, hexEdge.r - 1, BottomLeft}
		bottomLeftHexEdge = HexEdge{hexEdge.q, hexEdge.r - 1, BottomLeft}
		bottomRightHexEdge = HexEdge{hexEdge.q, hexEdge.r, MiddleLeft}
	case MiddleLeft:
		topLeftHexEdge = HexEdge{hexEdge.q, hexEdge.r - 1, BottomLeft}
		topRightHexEdge = HexEdge{hexEdge.q, hexEdge.r, TopLeft}
		bottomLeftHexEdge = HexEdge{hexEdge.q - 1, hexEdge.r + 1, TopLeft}
		bottomRightHexEdge = HexEdge{hexEdge.q, hexEdge.r, BottomLeft}
	case BottomLeft:
		topLeftHexEdge = HexEdge{hexEdge.q - 1, hexEdge.r + 1, TopLeft}
		topRightHexEdge = HexEdge{hexEdge.q, hexEdge.r, MiddleLeft}
		bottomLeftHexEdge = HexEdge{hexEdge.q, hexEdge.r + 1, MiddleLeft}
		bottomRightHexEdge = HexEdge{hexEdge.q, hexEdge.r + 1, TopLeft}
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
	location HexEdgeLocation
}

func (h HexEdge) GetQ() int {
	return h.q
}

func (h HexEdge) GetR() int {
	return h.r
}

func (h HexEdge) GetLocation() HexEdgeLocation {
	return h.location
}

func (h HexEdge) isAdjacentWithHexCorner(hexCorner HexCorner) bool {
	adjacentHexCorners := findAdjacentHexCornersFromHexEdge(h)

	return slices.Contains(adjacentHexCorners, hexCorner)
}

func (h HexEdge) isAdjacentWithHexEdge(hexEdge HexEdge) bool {
	adjacentHexEdges := findAdjacentHexEdgesFromHexEdge(h)

	return slices.Contains(adjacentHexEdges, hexEdge)
}
