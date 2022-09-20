package models

import (
	"github.com/pkg/errors"
)

func NewConstructionType(enum string) (ConstructionType, error) {
	constructionType := ConstructionType{enum}

	switch constructionType {
	case Settlement, City:
		return constructionType, nil
	default:
		return constructionType, errors.New("construction type is invalid")
	}
}

type ConstructionType struct {
	enum string
}

func (c ConstructionType) String() string {
	return c.enum
}

var (
	Settlement = ConstructionType{"SETTLEMENT"}
	City       = ConstructionType{"CITY"}
)
