package models

type hexDirection struct {
	q int
	r int
}

func (h hexDirection) multiply(scalar int) hexDirection {
	return hexDirection{
		q: h.q * scalar,
		r: h.r * scalar,
	}
}

func (h hexDirection) calculateEndpoint(startPoint Hex) Hex {
	return Hex{
		q: startPoint.q + h.q,
		r: startPoint.r + h.r,
	}
}
