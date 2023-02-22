package models

type aggregateRoot struct {
	aggregate
	events  []interface{}
	version int
}

func (a aggregateRoot) GetEvents() interface{} {
	return a.events
}

func (a aggregateRoot) GetVersion() int {
	return a.version
}
