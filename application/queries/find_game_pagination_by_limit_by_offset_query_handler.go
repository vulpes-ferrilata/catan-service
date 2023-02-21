package queries

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/catan-service/view/projectors"
)

type FindGamePaginationByLimitByOffsetQuery struct {
	Limit  int
	Offset int
}

func NewFindGamePaginationByLimitByOffsetQueryHandler(gamePaginationProjector projectors.GamePaginationProjector) *FindGamePaginationByLimitByOffsetQueryHandler {
	return &FindGamePaginationByLimitByOffsetQueryHandler{
		gamePaginationProjector: gamePaginationProjector,
	}
}

type FindGamePaginationByLimitByOffsetQueryHandler struct {
	gamePaginationProjector projectors.GamePaginationProjector
}

func (f FindGamePaginationByLimitByOffsetQueryHandler) Handle(ctx context.Context, findGamePaginationByLimitByOffsetQuery *FindGamePaginationByLimitByOffsetQuery) (*models.Pagination[*models.Game], error) {
	gamePagination, err := f.gamePaginationProjector.FindByLimitByOffset(ctx, findGamePaginationByLimitByOffsetQuery.Limit, findGamePaginationByLimitByOffsetQuery.Offset)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gamePagination, nil
}
