package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func ToGameView(gameDocument *documents.Game) *models.Game {
	if gameDocument == nil {
		return nil
	}

	activePlayer := toPlayerView(gameDocument.ActivePlayer)

	players, _ := slices.Map(func(playerDocument *documents.Player) (*models.Player, error) {
		return toPlayerView(playerDocument), nil
	}, gameDocument.Players)

	dices, _ := slices.Map(func(diceDocument *documents.Dice) (*models.Dice, error) {
		return toDiceView(diceDocument), nil
	}, gameDocument.Dices)

	achievements, _ := slices.Map(func(achievementDocument *documents.Achievement) (*models.Achievement, error) {
		return toAchievementView(achievementDocument), nil
	}, gameDocument.Achievements)

	resourceCards, _ := slices.Map(func(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
		return toResourceCardView(resourceCardDocument), nil
	}, gameDocument.ResourceCards)

	developmentCards, _ := slices.Map(func(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
		return toDevelopmentCardView(developmentCardDocument), nil
	}, gameDocument.DevelopmentCards)

	terrains, _ := slices.Map(func(terrainDocument *documents.Terrain) (*models.Terrain, error) {
		return toTerrainView(terrainDocument), nil
	}, gameDocument.Terrains)

	lands, _ := slices.Map(func(landDocument *documents.Land) (*models.Land, error) {
		return toLandView(landDocument), nil
	}, gameDocument.Lands)

	paths, _ := slices.Map(func(pathDocument *documents.Path) (*models.Path, error) {
		return toPathView(pathDocument), nil
	}, gameDocument.Paths)

	return &models.Game{
		ID:               gameDocument.ID,
		Status:           gameDocument.Status,
		Phase:            gameDocument.Phase,
		Turn:             gameDocument.Turn,
		ActivePlayer:     activePlayer,
		Players:          players,
		Dices:            dices,
		Achievements:     achievements,
		ResourceCards:    resourceCards,
		DevelopmentCards: developmentCards,
		Terrains:         terrains,
		Lands:            lands,
		Paths:            paths,
	}
}
