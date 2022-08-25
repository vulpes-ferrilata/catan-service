package models

import (
	"github.com/pkg/errors"
)

func NewConstructionType(value string) (constructionType, error) {
	constructionType := constructionType(value)

	switch constructionType {
	case Settlement, City:
		return constructionType, nil
	default:
		return constructionType, errors.New("construction type is invalid")
	}
}

type constructionType string

const (
	Settlement constructionType = "SETTLEMENT"
	City       constructionType = "CITY"
)
