package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
)

func ToGameDocument(game *models.Game) *documents.Game {
	activePlayerDocument := toPlayerDocument(game.GetActivePlayer())

	playerDocuments, _ := slices.Map(func(player *models.Player) (*documents.Player, error) {
		return toPlayerDocument(player), nil
	}, game.GetPlayers())

	diceDocuments, _ := slices.Map(func(dice *models.Dice) (*documents.Dice, error) {
		return toDiceDocument(dice), nil
	}, game.GetDices())

	achievementDocuments, _ := slices.Map(func(achievement *models.Achievement) (*documents.Achievement, error) {
		return toAchievementDocument(achievement), nil
	}, game.GetAchievements())

	resourceCardDocuments, _ := slices.Map(func(resourceCard *models.ResourceCard) (*documents.ResourceCard, error) {
		return toResourceCardDocument(resourceCard), nil
	}, game.GetResourceCards())

	developmentCardDocuments, _ := slices.Map(func(developmentCard *models.DevelopmentCard) (*documents.DevelopmentCard, error) {
		return toDevelopmentCardDocument(developmentCard), nil
	}, game.GetDevelopmentCards())

	terrainDocuments, _ := slices.Map(func(terrain *models.Terrain) (*documents.Terrain, error) {
		return toTerrainDocument(terrain), nil
	}, game.GetTerrains())

	landDocuments, _ := slices.Map(func(land *models.Land) (*documents.Land, error) {
		return toLandDocument(land), nil
	}, game.GetLands())

	pathDocuments, _ := slices.Map(func(path *models.Path) (*documents.Path, error) {
		return toPathDocument(path), nil
	}, game.GetPaths())

	return &documents.Game{
		DocumentRoot: documents.DocumentRoot{
			Document: documents.Document{
				ID: game.GetID(),
			},
			Version: game.GetVersion(),
		},
		Status:           game.GetStatus().String(),
		Phase:            game.GetPhase().String(),
		Turn:             game.GetTurn(),
		ActivePlayer:     activePlayerDocument,
		Players:          playerDocuments,
		Dices:            diceDocuments,
		Achievements:     achievementDocuments,
		ResourceCards:    resourceCardDocuments,
		DevelopmentCards: developmentCardDocuments,
		Terrains:         terrainDocuments,
		Lands:            landDocuments,
		Paths:            pathDocuments,
	}
}

func ToGameDomain(gameDocument *documents.Game) (*models.Game, error) {
	if gameDocument == nil {
		return nil, nil
	}

	status, err := models.NewGameStatus(gameDocument.Status)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	phase, err := models.NewGamePhase(gameDocument.Phase)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	activePlayer, err := toPlayerDomain(gameDocument.ActivePlayer)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	players, err := slices.Map(func(playerDocument *documents.Player) (*models.Player, error) {
		return toPlayerDomain(playerDocument)
	}, gameDocument.Players)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dices, _ := slices.Map(func(diceDocument *documents.Dice) (*models.Dice, error) {
		return toDiceDomain(diceDocument), nil
	}, gameDocument.Dices)

	achievements, err := slices.Map(func(achievementDocument *documents.Achievement) (*models.Achievement, error) {
		return toAchievementDomain(achievementDocument)
	}, gameDocument.Achievements)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCards, err := slices.Map(func(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
		return toResourceCardDomain(resourceCardDocument)
	}, gameDocument.ResourceCards)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCards, err := slices.Map(func(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
		return toDevelopmentCardDomain(developmentCardDocument)
	}, gameDocument.DevelopmentCards)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	terrains, err := slices.Map(func(terrainDocument *documents.Terrain) (*models.Terrain, error) {
		return toTerrainDomain(terrainDocument)
	}, gameDocument.Terrains)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	lands, err := slices.Map(func(landDocument *documents.Land) (*models.Land, error) {
		return toLandDomain(landDocument)
	}, gameDocument.Lands)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	paths, err := slices.Map(func(pathDocument *documents.Path) (*models.Path, error) {
		return toPathDomain(pathDocument)
	}, gameDocument.Paths)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	game := models.GameBuilder{}.
		SetID(gameDocument.ID).
		SetStatus(status).
		SetPhase(phase).
		SetTurn(gameDocument.Turn).
		SetActivePlayer(activePlayer).
		SetPlayers(players...).
		SetDices(dices...).
		SetAchievements(achievements...).
		SetResourceCards(resourceCards...).
		SetDevelopmentCards(developmentCards...).
		SetTerrains(terrains...).
		SetLands(lands...).
		SetPaths(paths...).
		SetVersion(gameDocument.Version).
		Create()

	return game, nil
}
