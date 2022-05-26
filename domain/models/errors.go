package models

import "github.com/pkg/errors"

var (
	ErrEnumIsInvalid error = errors.New("invalid enum")
)
