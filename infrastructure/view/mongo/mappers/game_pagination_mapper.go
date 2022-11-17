package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type GamePaginationMapper struct{}

func (g GamePaginationMapper) ToView(gamePaginationDocument *documents.Pagination[*documents.Game]) (*models.Pagination[*models.Game], error) {
	if gamePaginationDocument == nil {
		return nil, nil
	}

	games, err := slices.Map(func(gameDocument *documents.Game) (*models.Game, error) {
		return gameMapper{}.ToView(gameDocument)
	}, gamePaginationDocument.Data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.Pagination[*models.Game]{
		Total: gamePaginationDocument.Metadata.Total,
		Data:  games,
	}, nil
}
