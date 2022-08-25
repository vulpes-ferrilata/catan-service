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

	harborResponses, _ := slices.Map(func(harbor *models.Harbor) (*catan.HarborResponse, error) {
		return toHarborResponse(harbor), nil
	}, game.Harbors)

	robberResponse := toRobberResponse(game.Robber)

	landResponses, _ := slices.Map(func(land *models.Land) (*catan.LandResponse, error) {
		return toLandResponse(land), nil
	}, game.Lands)

	pathResponses, _ := slices.Map(func(path *models.Path) (*catan.PathResponse, error) {
		return toPathResponse(path), nil
	}, game.Paths)

	return &catan.GameResponse{
		ID:               game.ID.Hex(),
		Status:           game.Status,
		Turn:             int32(game.Turn),
		IsRolledDices:    game.IsRolledDices,
		Players:          playerResponses,
		Dices:            diceResponses,
		Achievements:     achievementResponses,
		ResourceCards:    resourceCardResponses,
		DevelopmentCards: developmentCardResponses,
		Terrains:         terrainResponses,
		Harbors:          harborResponses,
		Robber:           robberResponse,
		Lands:            landResponses,
		Paths:            pathResponses,
	}
}
