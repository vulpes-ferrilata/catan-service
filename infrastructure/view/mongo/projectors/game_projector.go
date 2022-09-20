package projectors

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/mappers"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/catan-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewGameProjector(db *mongo.Database) projectors.GameProjector {
	return &gameProjector{
		gameCollection: db.Collection("games"),
	}
}

type gameProjector struct {
	gameCollection *mongo.Collection
}

func (g gameProjector) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.Game, error) {
	gameDocuments := make([]*documents.Game, 0)

	gameCursor, err := g.gameCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := gameCursor.All(ctx, &gameDocuments); err != nil {
		return nil, errors.WithStack(err)
	}

	games, _ := slices.Map(func(gameDocument *documents.Game) (*models.Game, error) {
		return mappers.ToGameDetailView(gameDocument), nil
	}, gameDocuments)

	for _, game := range games {
		for _, player := range game.Players {
			if player.UserID != userID && game.Status == "STARTED" {
				for _, resourceCard := range player.ResourceCards {
					if !resourceCard.IsSelected {
						resourceCard.Type = "HIDDEN"
					}
				}

				for _, developmentCard := range player.DevelopmentCards {
					if developmentCard.Status != "USED" {
						developmentCard.Type = "HIDDEN"
					}
				}

				player.Score = 0
			}

			if player.UserID == userID {
				game.Me = player
				game.Players = slices.Remove(game.Players, player)
			}
		}
	}

	return games, nil
}

func (g gameProjector) GetByIDByUserID(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*models.Game, error) {
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

	game := mappers.ToGameDetailView(gameDocument)

	for _, player := range game.Players {
		if player.UserID != userID && game.Status == "STARTED" {
			for _, resourceCard := range player.ResourceCards {
				if !resourceCard.IsSelected {
					resourceCard.Type = "HIDDEN"
				}
			}

			for _, developmentCard := range player.DevelopmentCards {
				if developmentCard.Status != "USED" {
					developmentCard.Type = "HIDDEN"
				}
			}

			player.Score = 0
		}

		if player.UserID == userID {
			game.Me = player
			game.Players = slices.Remove(game.Players, player)
		}
	}

	return game, nil
}
