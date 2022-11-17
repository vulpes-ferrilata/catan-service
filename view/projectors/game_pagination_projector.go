package projectors

import (
	"context"

	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type GamePaginationProjector interface {
	FindByLimitByOffset(ctx context.Context, limit int, offset int) (*models.Pagination[*models.Game], error)
}
