package models

type Harbor struct {
	aggregate
	hex        Hex
	harborType HarborType
}

func (t Harbor) GetHex() Hex {
	return t.hex
}

func (t Harbor) GetType() HarborType {
	return t.harborType
}
