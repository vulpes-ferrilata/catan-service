package mappers

import (
	"github.com/pkg/errors"
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/slices"
)

type GameDetailMapper struct{}

func (g GameDetailMapper) ToResponse(gameDetail *models.GameDetail) (*pb_models.GameDetail, error) {
	if gameDetail == nil {
		return nil, nil
	}

	activePlayerResponse, err := playerMapper{}.ToResponse(gameDetail.ActivePlayer)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	playerResponses, err := slices.Map(func(player *models.Player) (*pb_models.Player, error) {
		return playerMapper{}.ToResponse(player)
	}, gameDetail.Players...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	diceResponses, err := slices.Map(func(dice *models.Dice) (*pb_models.Dice, error) {
		return diceMapper{}.ToResponse(dice)
	}, gameDetail.Dices...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievementResponses, err := slices.Map(func(achievement *models.Achievement) (*pb_models.Achievement, error) {
		return achievementMapper{}.ToResponse(achievement)
	}, gameDetail.Achievements...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCardResponses, err := slices.Map(func(resourceCard *models.ResourceCard) (*pb_models.ResourceCard, error) {
		return resourceCardMapper{}.ToResponse(resourceCard)
	}, gameDetail.ResourceCards...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCardResponses, err := slices.Map(func(developmentCard *models.DevelopmentCard) (*pb_models.DevelopmentCard, error) {
		return developmentCardMapper{}.ToResponse(developmentCard)
	}, gameDetail.DevelopmentCards...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	terrainResponses, err := slices.Map(func(terrain *models.Terrain) (*pb_models.Terrain, error) {
		return terrainMapper{}.ToResponse(terrain)
	}, gameDetail.Terrains...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	landResponses, err := slices.Map(func(land *models.Land) (*pb_models.Land, error) {
		return landMapper{}.ToResponse(land)
	}, gameDetail.Lands...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pathResponses, err := slices.Map(func(path *models.Path) (*pb_models.Path, error) {
		return pathMapper{}.ToResponse(path)
	}, gameDetail.Paths...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb_models.GameDetail{
		ID:               gameDetail.ID.Hex(),
		Status:           gameDetail.Status,
		Phase:            gameDetail.Phase,
		Turn:             int32(gameDetail.Turn),
		ActivePlayer:     activePlayerResponse,
		Players:          playerResponses,
		Dices:            diceResponses,
		Achievements:     achievementResponses,
		ResourceCards:    resourceCardResponses,
		DevelopmentCards: developmentCardResponses,
		Terrains:         terrainResponses,
		Lands:            landResponses,
		Paths:            pathResponses,
	}, nil
}
