package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DiceBuilder struct {
	id     primitive.ObjectID
	number int
}

func (a DiceBuilder) SetID(id primitive.ObjectID) DiceBuilder {
	a.id = id

	return a
}

func (a DiceBuilder) SetNumber(number int) DiceBuilder {
	a.number = number

	return a
}

func (a DiceBuilder) Create() *Dice {
	return &Dice{
		aggregate: aggregate{
			id: a.id,
		},
		number: a.number,
	}
}
