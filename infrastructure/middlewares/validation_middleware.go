package middlewares

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/infrastructure/bus"
	app_errors "github.com/VulpesFerrilata/catan-service/infrastructure/errors"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func NewValidationMiddleware(validate *validator.Validate) *ValidationMiddleware {
	return &ValidationMiddleware{
		validate: validate,
	}
}

type ValidationMiddleware struct {
	validate *validator.Validate
}

func (v ValidationMiddleware) WrapHandler(next bus.HandlerFunc) bus.HandlerFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		err := v.validate.StructCtx(ctx, input)
		if validationErrs, ok := errors.Cause(err).(validator.ValidationErrors); ok {
			return nil, app_errors.NewValidationError(validationErrs...)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return next(ctx, input)
	}
}
