package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/application/queries"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/query/wrappers"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/catan-service/view/projectors"
)

func NewFindGamePaginationByLimitByOffsetQueryHandler(validate *validator.Validate, gamePaginationProjector projectors.GamePaginationProjector) query.QueryHandler[*queries.FindGamePaginationByLimitByOffset, *models.Pagination[*models.Game]] {
	handler := &findGamePaginationByLimitByOffsetQueryHandler{
		gamePaginationProjector: gamePaginationProjector,
	}
	validationWrapper := wrappers.NewValidationWrapper[*queries.FindGamePaginationByLimitByOffset, *models.Pagination[*models.Game]](validate, handler)

	return validationWrapper
}

type findGamePaginationByLimitByOffsetQueryHandler struct {
	gamePaginationProjector projectors.GamePaginationProjector
}

func (f findGamePaginationByLimitByOffsetQueryHandler) Handle(ctx context.Context, findGamePaginationByLimitByOffsetQuery *queries.FindGamePaginationByLimitByOffset) (*models.Pagination[*models.Game], error) {
	gamePagination, err := f.gamePaginationProjector.FindByLimitByOffset(ctx, findGamePaginationByLimitByOffsetQuery.Limit, findGamePaginationByLimitByOffsetQuery.Offset)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gamePagination, nil
}
