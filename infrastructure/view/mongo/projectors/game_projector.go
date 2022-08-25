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

// func (g gameProjector) getProjectionPipeline(userID primitive.ObjectID) mongo.Pipeline {

// 	return mongo.Pipeline{
// 		//separate mes with players
// 		{{"$project", bson.D{
// 			{"_id", 1},
// 			{"status", 1},
// 			{"turn", 1},
// 			{"is_rolled_dices", 1},
// 			{"mes", bson.D{
// 				{"$filter", bson.D{
// 					{"input", "$players"},
// 					{"cond", bson.D{
// 						{"$eq", bson.A{
// 							"$$this.user_id",
// 							userID,
// 						}},
// 					}},
// 				}},
// 			}},
// 			{"players", bson.D{
// 				{"$filter", bson.D{
// 					{"input", "$players"},
// 					{"cond", bson.D{
// 						{"$ne", bson.A{
// 							"$$this.user_id",
// 							userID,
// 						}},
// 					}},
// 				}},
// 			}},
// 			{"dices", 1},
// 			{"achievements", 1},
// 			{"resource_cards", 1},
// 			{"development_cards", 1},
// 			{"terrains", 1},
// 			{"harbors", 1},
// 			{"robber", 1},
// 			{"lands", 1},
// 			{"paths", 1},
// 		}}},
// 		//add fields is_me to mes
// 		{{"$addFields", bson.D{
// 			{"mes", bson.D{
// 				{"$map", bson.D{
// 					{"input", "$mes"},
// 					{"in", bson.D{
// 						{"_id", "$$this._id"},
// 						{"user_id", "$$this.user_id"},
// 						{"turn_order", "$$this.turn_order"},
// 						{"color", "$$this.color"},
// 						{"is_active", "$$this.is_active"},
// 						{"is_trading", "$$this.is_trading"},
// 						{"is_me", true},
// 						{"achievements", "$$this.achievements"},
// 						{"resource_cards", "$$this.resource_cards"},
// 						{"development_cards", "$$this.development_cards"},
// 						{"constructions", "$$this.constructions"},
// 						{"roads", "$$this.roads"},
// 					}},
// 				}},
// 			}},
// 		}}},
// 		//add fields is_me to players
// 		{{"$addFields", bson.D{
// 			{"players", bson.D{
// 				{"$map", bson.D{
// 					{"input", "$players"},
// 					{"in", bson.D{
// 						{"_id", "$$this._id"},
// 						{"user_id", "$$this.user_id"},
// 						{"turn_order", "$$this.turn_order"},
// 						{"color", "$$this.color"},
// 						{"is_active", "$$this.is_active"},
// 						{"is_trading", "$$this.is_trading"},
// 						{"is_me", false},
// 						{"achievements", "$$this.achievements"},
// 						{"resource_cards", "$$this.resource_cards"},
// 						{"development_cards", "$$this.development_cards"},
// 						{"constructions", "$$this.constructions"},
// 						{"roads", "$$this.roads"},
// 					}},
// 				}},
// 			}},
// 		}}},
// 		//concat me with players
// 		{{"$project", bson.D{
// 			{"_id", 1},
// 			{"status", 1},
// 			{"turn", 1},
// 			{"is_rolled_dices", 1},
// 			{"players", bson.D{
// 				{"$concatArrays", bson.A{
// 					"$mes",
// 					"$players",
// 				},
// 				},
// 			},
// 			},
// 			{"dices", 1},
// 			{"achievements", 1},
// 			{"resource_cards", 1},
// 			{"development_cards", 1},
// 			{"terrains", 1},
// 			{"harbors", 1},
// 			{"robber", 1},
// 			{"lands", 1},
// 			{"paths", 1},
// 		}}},
// 	}
// }

func (g gameProjector) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.Game, error) {
	gameDocuments := make([]*documents.Game, 0)

	// projectionPipeline := g.getProjectionPipeline(userID)

	// pipeline := mongo.Pipeline{}

	// pipeline = append(pipeline, projectionPipeline...)

	// gameCursor, err := g.gameCollection.Aggregate(ctx, pipeline)
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }

	// if err := gameCursor.All(ctx, &gameDocuments); err != nil {
	// 	return nil, errors.WithStack(err)
	// }

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
			if player.UserID != userID {
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

				continue
			}

			player.IsMe = true
		}
	}

	return games, nil
}

func (g gameProjector) GetByIDByUserID(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*models.Game, error) {
	gameDocument := &documents.Game{}

	// filterStage := bson.D{
	// 	{"$match", bson.D{{"_id", id}}},
	// }

	// projectionPipeline := g.getProjectionPipeline(userID)

	// pipeline := mongo.Pipeline{}

	// pipeline = append(pipeline, filterStage)
	// pipeline = append(pipeline, projectionPipeline...)

	// gameCursor, err := g.gameCollection.Aggregate(ctx, pipeline)
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }

	// if isExists := gameCursor.Next(ctx); !isExists {
	// 	return nil, errors.WithStack(app_errors.ErrGameNotFound)
	// }

	// if err := gameCursor.Decode(gameDocument); err != nil {
	// 	return nil, errors.WithStack(err)
	// }

	// defer gameCursor.Close(ctx)
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
		if player.UserID != userID {
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

			continue
		}

		player.IsMe = true
	}

	return game, nil
}
