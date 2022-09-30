package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func ToGameResponse(game *models.Game) *responses.Game {
	if game == nil {
		return nil
	}

	activePlayerResponse := toPlayerResponse(game.ActivePlayer)

	playerResponses, _ := slices.Map(func(player *models.Player) (*responses.Player, error) {
		return toPlayerResponse(player), nil
	}, game.Players)

	diceResponses, _ := slices.Map(func(dice *models.Dice) (*responses.Dice, error) {
		return toDiceResponse(dice), nil
	}, game.Dices)

	achievementResponses, _ := slices.Map(func(achievement *models.Achievement) (*responses.Achievement, error) {
		return toAchievementResponse(achievement), nil
	}, game.Achievements)

	resourceCardResponses, _ := slices.Map(func(resourceCard *models.ResourceCard) (*responses.ResourceCard, error) {
		return toResourceCardResponse(resourceCard), nil
	}, game.ResourceCards)

	developmentCardResponses, _ := slices.Map(func(developmentCard *models.DevelopmentCard) (*responses.DevelopmentCard, error) {
		return toDevelopmentCardResponse(developmentCard), nil
	}, game.DevelopmentCards)

	terrainResponses, _ := slices.Map(func(terrain *models.Terrain) (*responses.Terrain, error) {
		return toTerrainResponse(terrain), nil
	}, game.Terrains)

	landResponses, _ := slices.Map(func(land *models.Land) (*responses.Land, error) {
		return toLandResponse(land), nil
	}, game.Lands)

	pathResponses, _ := slices.Map(func(path *models.Path) (*responses.Path, error) {
		return toPathResponse(path), nil
	}, game.Paths)

	return &responses.Game{
		ID:               game.ID.Hex(),
		Status:           game.Status,
		Phase:            game.Phase,
		Turn:             int32(game.Turn),
		ActivePlayer:     activePlayerResponse,
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
