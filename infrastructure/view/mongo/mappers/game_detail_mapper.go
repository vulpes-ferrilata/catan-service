package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type GameDetailMapper struct{}

func (g GameDetailMapper) ToView(gameDetailDocument *documents.GameDetail) (*models.GameDetail, error) {
	if gameDetailDocument == nil {
		return nil, nil
	}

	activePlayer, err := playerMapper{}.ToView(gameDetailDocument.ActivePlayer)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	players, err := slices.Map(func(playerDocument *documents.Player) (*models.Player, error) {
		return playerMapper{}.ToView(playerDocument)
	}, gameDetailDocument.Players)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dices, err := slices.Map(func(diceDocument *documents.Dice) (*models.Dice, error) {
		return diceMapper{}.ToView(diceDocument)
	}, gameDetailDocument.Dices)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievements, err := slices.Map(func(achievementDocument *documents.Achievement) (*models.Achievement, error) {
		return achievementMapper{}.ToView(achievementDocument)
	}, gameDetailDocument.Achievements)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCards, err := slices.Map(func(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
		return resourceCardMapper{}.ToView(resourceCardDocument)
	}, gameDetailDocument.ResourceCards)

	developmentCards, err := slices.Map(func(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
		return developmentCardMapper{}.ToView(developmentCardDocument)
	}, gameDetailDocument.DevelopmentCards)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	terrains, err := slices.Map(func(terrainDocument *documents.Terrain) (*models.Terrain, error) {
		return terrainMapper{}.ToView(terrainDocument)
	}, gameDetailDocument.Terrains)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	lands, err := slices.Map(func(landDocument *documents.Land) (*models.Land, error) {
		return landMapper{}.ToView(landDocument)
	}, gameDetailDocument.Lands)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	paths, err := slices.Map(func(pathDocument *documents.Path) (*models.Path, error) {
		return pathMapper{}.ToView(pathDocument)
	}, gameDetailDocument.Paths)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.GameDetail{
		ID:               gameDetailDocument.ID,
		Status:           gameDetailDocument.Status,
		Phase:            gameDetailDocument.Phase,
		Turn:             gameDetailDocument.Turn,
		ActivePlayer:     activePlayer,
		Players:          players,
		Dices:            dices,
		Achievements:     achievements,
		ResourceCards:    resourceCards,
		DevelopmentCards: developmentCards,
		Terrains:         terrains,
		Lands:            lands,
		Paths:            paths,
	}, nil
}
