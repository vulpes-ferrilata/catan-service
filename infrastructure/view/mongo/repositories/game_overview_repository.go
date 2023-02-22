package repositories

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/mappers"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/catan-service/view/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewGameOverviewRepository(db *mongo.Database) repositories.GameOverviewRepository {
	return &gameOverviewRepository{
		gameOverviewCollection: db.Collection("game_overviews"),
	}
}

type gameOverviewRepository struct {
	gameOverviewCollection *mongo.Collection
}

func (g gameOverviewRepository) InsertOrUpdate(ctx context.Context, gameOverview *models.GameOverview) error {
	gameOverviewDocument, err := mappers.GameOverviewMapper{}.ToDocument(gameOverview)
	if err != nil {
		return errors.WithStack(err)
	}

	filter := bson.M{
		"_id": gameOverviewDocument.ID,
	}

	if _, err := g.gameOverviewCollection.ReplaceOne(ctx, filter, gameOverviewDocument, options.Replace().SetUpsert(true)); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
