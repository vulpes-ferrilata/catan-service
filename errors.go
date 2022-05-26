package service

import (
	app_errors "github.com/VulpesFerrilata/catan-service/infrastructure/errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/pkg/errors"
)

var (
	ErrHeaderFormatIsInvalid = NewDetailError("invalid-header-format")
	ErrPlayerIsFull          = NewDetailError("full-player")
	ErrPlayerAlreadyJoined   = NewDetailError("joined-player")
	ErrOtherPlayerTurn       = NewDetailError("other-player-turn")
)

func NewDetailError(translationKey string) app_errors.DetailError {
	return &detailError{
		translationKey: translationKey,
	}
}

type detailError struct {
	translationKey string
}

func (d detailError) Error() string {
	return d.translationKey
}

func (d detailError) Translate(translator ut.Translator) (string, error) {
	message, err := translator.T(d.translationKey)
	if err != nil {
		return "", errors.Wrap(err, d.translationKey)
	}

	return message, nil
}
