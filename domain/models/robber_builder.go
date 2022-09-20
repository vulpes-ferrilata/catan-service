package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RobberBuilder struct {
	id primitive.ObjectID
}

func (r RobberBuilder) SetID(id primitive.ObjectID) RobberBuilder {
	r.id = id

	return r
}

func (r RobberBuilder) Create() *Robber {
	return &Robber{
		aggregate: aggregate{
			id: r.id,
		},
	}
}
