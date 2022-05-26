package models

import (
	"github.com/VulpesFerrilata/catan-service/domain/models/common"
	"github.com/google/uuid"
)

func NewDice(id uuid.UUID, number int) *Dice {
	return &Dice{
		Entity: common.NewEntity(id),
		number: number,
	}
}

type Dice struct {
	common.Entity
	number int
}

func (d Dice) GetNumber() int {
	return d.number
}
