package models

import (
	"github.com/pkg/errors"
)

func NewDevelopmentCardStatus(value string) (developmentCardStatus, error) {
	developmentCardStatus := developmentCardStatus(value)

	switch developmentCardStatus {
	case Enable, Disable, Used:
		return developmentCardStatus, nil
	default:
		return developmentCardStatus, errors.New("development card status is invalid")
	}
}

type developmentCardStatus string

const (
	Enable  developmentCardStatus = "ENABLE"
	Disable developmentCardStatus = "DISABLE"
	Used    developmentCardStatus = "USED"
)
