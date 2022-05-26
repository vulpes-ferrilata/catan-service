package infrastructure

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
)

var (
	ErrTranslatorNotFound = errors.New("translator not found")
)

func NewValidate(universalTranslator *ut.UniversalTranslator) (*validator.Validate, error) {
	validate := validator.New()

	en := en.New()

	translator, ok := universalTranslator.GetTranslator(en.Locale())
	if !ok {
		return nil, errors.Wrap(ErrTranslatorNotFound, en.Locale())
	}

	if err := en_translations.RegisterDefaultTranslations(validate, translator); err != nil {
		return nil, errors.WithStack(err)
	}

	return validate, nil
}
