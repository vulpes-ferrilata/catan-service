package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoadBuilder interface {
	SetID(id primitive.ObjectID) RoadBuilder
	SetPath(path *Path) RoadBuilder
	Create() *Road
}

func NewRoadBuilder() RoadBuilder {
	return &roadBuilder{}
}

type roadBuilder struct {
	id   primitive.ObjectID
	path *Path
}

func (r *roadBuilder) SetID(id primitive.ObjectID) RoadBuilder {
	r.id = id

	return r
}

func (r *roadBuilder) SetPath(path *Path) RoadBuilder {
	r.path = path

	return r
}

func (r *roadBuilder) Create() *Road {
	return &Road{
		aggregate: aggregate{
			id: r.id,
		},
		path: r.path,
	}
}
