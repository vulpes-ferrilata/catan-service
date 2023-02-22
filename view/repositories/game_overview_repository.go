package repositories

import (
	"context"

	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type GameOverviewRepository interface {
	InsertOrUpdate(ctx context.Context, gameOverview *models.GameOverview) error
}
