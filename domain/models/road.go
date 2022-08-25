package models

type Road struct {
	aggregate
	path *Path
}

func (r Road) GetPath() *Path {
	return r.path
}
