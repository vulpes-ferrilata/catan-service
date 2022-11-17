package models

type Pagination[Model any] struct {
	Total int
	Data  []Model
}
