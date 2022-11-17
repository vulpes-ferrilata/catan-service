package repositories

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/mappers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewGameRepository(db *mongo.Database) repositories.GameRepository {
	return &gameRepository{
		gameCollection: db.Collection("games"),
	}
}

type gameRepository struct {
	gameCollection *mongo.Collection
}

func (g gameRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Game, error) {
	gameDocument := &documents.Game{}

	filter := bson.M{
		"_id": id,
	}

	err := g.gameCollection.FindOne(ctx, filter).Decode(gameDocument)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, errors.WithStack(app_errors.ErrGameNotFound)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	game, err := mappers.GameMapper{}.ToDomain(gameDocument)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return game, nil
}

func (g gameRepository) Insert(ctx context.Context, game *models.Game) error {
	gameDocument, err := mappers.GameMapper{}.ToDocument(game)
	if err != nil {
		return errors.WithStack(err)
	}

	gameDocument.Version = 1

	if _, err := g.gameCollection.InsertOne(ctx, gameDocument); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g gameRepository) Update(ctx context.Context, game *models.Game) error {
	gameDocument, err := mappers.GameMapper{}.ToDocument(game)
	if err != nil {
		return errors.WithStack(err)
	}

	filter := bson.M{
		"_id":     gameDocument.ID,
		"version": gameDocument.Version,
	}

	gameDocument.Version++

	result, err := g.gameCollection.ReplaceOne(ctx, filter, gameDocument)
	if err != nil {
		return errors.WithStack(err)
	}
	if result.ModifiedCount == 0 {
		return errors.WithStack(app_errors.ErrStaleGame)
	}

	return nil
}
