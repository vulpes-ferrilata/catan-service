package models

func NewHexDirection(q int, r int) hexDirection {
	return hexDirection{
		q: q,
		r: r,
	}
}

type hexDirection struct {
	q int
	r int
}

func (h hexDirection) GetQ() int {
	return h.q
}

func (h hexDirection) GetR() int {
	return h.r
}

func (h hexDirection) Multiply(scalar int) hexDirection {
	return hexDirection{
		q: h.GetQ() * scalar,
		r: h.GetR() * scalar,
	}
}

func (h hexDirection) CalculateEndpoint(startPoint Hex) Hex {
	return Hex{
		q: startPoint.GetQ() + h.GetQ(),
		r: startPoint.GetR() + h.GetR(),
	}
}
