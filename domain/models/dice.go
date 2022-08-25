package models

type Dice struct {
	aggregate
	number int
}

func (d Dice) GetNumber() int {
	return d.number
}
