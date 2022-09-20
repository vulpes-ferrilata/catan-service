package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoadBuilder struct {
	id   primitive.ObjectID
	path *Path
}

func (r RoadBuilder) SetID(id primitive.ObjectID) RoadBuilder {
	r.id = id

	return r
}

func (r RoadBuilder) SetPath(path *Path) RoadBuilder {
	r.path = path

	return r
}

func (r RoadBuilder) Create() *Road {
	return &Road{
		aggregate: aggregate{
			id: r.id,
		},
		path: r.path,
	}
}
