package mappers

import (
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/slices"
)

type GamePaginationMapper struct{}

func (g GamePaginationMapper) ToResponse(gamePagination *models.Pagination[*models.Game]) (*pb_models.GamePagination, error) {
	if gamePagination == nil {
		return nil, nil
	}

	games, _ := slices.Map(func(game *models.Game) (*pb_models.Game, error) {
		return gameMapper{}.ToResponse(game)
	}, gamePagination.Data...)

	return &pb_models.GamePagination{
		Total: int32(gamePagination.Total),
		Data:  games,
	}, nil
}
