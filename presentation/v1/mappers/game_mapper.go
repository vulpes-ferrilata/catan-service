package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func ToGameResponse(game *models.Game) *catan.GameResponse {
	if game == nil {
		return nil
	}

	me := toPlayerResponse(game.Me)

	playerResponses, _ := slices.Map(func(player *models.Player) (*catan.PlayerResponse, error) {
		return toPlayerResponse(player), nil
	}, game.Players)

	diceResponses, _ := slices.Map(func(dice *models.Dice) (*catan.DiceResponse, error) {
		return toDiceResponse(dice), nil
	}, game.Dices)

	achievementResponses, _ := slices.Map(func(achievement *models.Achievement) (*catan.AchievementResponse, error) {
		return toAchievementResponse(achievement), nil
	}, game.Achievements)

	resourceCardResponses, _ := slices.Map(func(resourceCard *models.ResourceCard) (*catan.ResourceCardResponse, error) {
		return toResourceCardResponse(resourceCard), nil
	}, game.ResourceCards)

	developmentCardResponses, _ := slices.Map(func(developmentCard *models.DevelopmentCard) (*catan.DevelopmentCardResponse, error) {
		return toDevelopmentCardResponse(developmentCard), nil
	}, game.DevelopmentCards)

	terrainResponses, _ := slices.Map(func(terrain *models.Terrain) (*catan.TerrainResponse, error) {
		return toTerrainResponse(terrain), nil
	}, game.Terrains)

	landResponses, _ := slices.Map(func(land *models.Land) (*catan.LandResponse, error) {
		return toLandResponse(land), nil
	}, game.Lands)

	pathResponses, _ := slices.Map(func(path *models.Path) (*catan.PathResponse, error) {
		return toPathResponse(path), nil
	}, game.Paths)

	return &catan.GameResponse{
		ID:               game.ID.Hex(),
		Status:           game.Status,
		Phase:            game.Phase,
		Turn:             int32(game.Turn),
		Me:               me,
		Players:          playerResponses,
		Dices:            diceResponses,
		Achievements:     achievementResponses,
		ResourceCards:    resourceCardResponses,
		DevelopmentCards: developmentCardResponses,
		Terrains:         terrainResponses,
		Lands:            landResponses,
		Paths:            pathResponses,
	}
}
