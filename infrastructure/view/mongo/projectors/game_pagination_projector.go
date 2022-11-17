package projectors

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/mappers"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/catan-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewGamePaginationProjector(db *mongo.Database) projectors.GamePaginationProjector {
	return &gamePaginationProjector{
		gameCollection: db.Collection("games"),
	}
}

type gamePaginationProjector struct {
	gameCollection *mongo.Collection
}

func (g gamePaginationProjector) FindByLimitByOffset(ctx context.Context, limit int, offset int) (*models.Pagination[*models.Game], error) {
	gamePaginationDocument := &documents.Pagination[*documents.Game]{}

	gameProjection := make(bson.A, 0)
	if offset > 0 {
		gameProjection = append(gameProjection, bson.D{{"$skip", offset}})
	}
	if limit > 0 {
		gameProjection = append(gameProjection, bson.D{{"$limit", limit}})
	}
	gameProjection = append(gameProjection, bson.D{
		{"$project",
			bson.D{
				{"player_quantity",
					bson.D{
						{"$sum",
							bson.A{
								bson.D{{"$size", "$players"}},
								1,
							},
						},
					},
				},
				{"status", 1},
			},
		},
	})

	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"status", bson.D{{"$ne", "FINISHED"}}}}}},
		bson.D{
			{"$facet",
				bson.D{
					{"metadata",
						bson.A{
							bson.D{
								{"$group",
									bson.D{
										{"_id", primitive.Null{}},
										{"total", bson.D{{"$sum", 1}}},
									},
								},
							},
						},
					},
					{"data", gameProjection},
				},
			},
		},
		bson.D{{"$unwind", "$metadata"}},
	}

	gameCursor, err := g.gameCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if gameCursor.TryNext(ctx) {
		if err := gameCursor.Decode(gamePaginationDocument); err != nil {
			return nil, errors.WithStack(err)
		}

		if err := gameCursor.Close(ctx); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	games, err := mappers.GamePaginationMapper{}.ToView(gamePaginationDocument)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return games, nil
}
