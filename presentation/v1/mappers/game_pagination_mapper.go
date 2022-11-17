package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type GamePaginationMapper struct{}

func (g GamePaginationMapper) ToResponse(gamePagination *models.Pagination[*models.Game]) (*responses.GamePagination, error) {
	if gamePagination == nil {
		return nil, nil
	}

	games, _ := slices.Map(func(game *models.Game) (*responses.Game, error) {
		return gameMapper{}.ToResponse(game)
	}, gamePagination.Data)

	return &responses.GamePagination{
		Total: int32(gamePagination.Total),
		Data:  games,
	}, nil
}
