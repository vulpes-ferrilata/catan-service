package projectors

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/app_errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/mappers"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/catan-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewGameDetailProjector(db *mongo.Database) projectors.GameDetailProjector {
	return &gameDetailProjector{
		gameCollection: db.Collection("games"),
	}
}

type gameDetailProjector struct {
	gameCollection *mongo.Collection
}

func (g gameDetailProjector) GetByIDByUserID(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*models.GameDetail, error) {
	gameDetailDocument := &documents.GameDetail{}

	filter := bson.M{
		"_id": id,
	}

	err := g.gameCollection.FindOne(ctx, filter).Decode(gameDetailDocument)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, errors.WithStack(app_errors.ErrGameNotFound)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if gameDetailDocument.Status == "Started" {
		for _, playerDocument := range append(gameDetailDocument.Players, gameDetailDocument.ActivePlayer) {
			if playerDocument.UserID != userID {
				for _, resourceCard := range playerDocument.ResourceCards {
					if !resourceCard.Offering {
						resourceCard.Type = "Hidden"
					}
				}

				for _, developmentCard := range playerDocument.DevelopmentCards {
					if developmentCard.Status != "Used" {
						developmentCard.Type = "Hidden"
					}
				}
			}
		}
	}

	gameDetail, err := mappers.GameDetailMapper{}.ToView(gameDetailDocument)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gameDetail, nil
}
