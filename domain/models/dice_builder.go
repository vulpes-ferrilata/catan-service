package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DiceBuilder interface {
	SetID(id primitive.ObjectID) DiceBuilder
	SetNumber(number int) DiceBuilder
	Create() *Dice
}

func NewDiceBuilder() DiceBuilder {
	return &diceBuilder{}
}

type diceBuilder struct {
	id     primitive.ObjectID
	number int
}

func (a *diceBuilder) SetID(id primitive.ObjectID) DiceBuilder {
	a.id = id

	return a
}

func (a *diceBuilder) SetNumber(number int) DiceBuilder {
	a.number = number

	return a
}

func (a diceBuilder) Create() *Dice {
	return &Dice{
		aggregate: aggregate{
			id: a.id,
		},
		number: a.number,
	}
}
