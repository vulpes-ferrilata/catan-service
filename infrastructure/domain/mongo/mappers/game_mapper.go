package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
	"github.com/vulpes-ferrilata/slices"
)

type GameMapper struct{}

func (g GameMapper) ToDocument(game *models.Game) (*documents.Game, error) {
	activePlayerDocument, err := playerMapper{}.ToDocument(game.GetActivePlayer())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	playerDocuments, err := slices.Map(func(player *models.Player) (*documents.Player, error) {
		return playerMapper{}.ToDocument(player)
	}, game.GetPlayers()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	diceDocuments, err := slices.Map(func(dice *models.Dice) (*documents.Dice, error) {
		return diceMapper{}.ToDocument(dice)
	}, game.GetDices()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievementDocuments, err := slices.Map(func(achievement *models.Achievement) (*documents.Achievement, error) {
		return achievementMapper{}.ToDocument(achievement)
	}, game.GetAchievements()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCardDocuments, err := slices.Map(func(resourceCard *models.ResourceCard) (*documents.ResourceCard, error) {
		return resourceCardMapper{}.ToDocument(resourceCard)
	}, game.GetResourceCards()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCardDocuments, err := slices.Map(func(developmentCard *models.DevelopmentCard) (*documents.DevelopmentCard, error) {
		return developmentCardMapper{}.ToDocument(developmentCard)
	}, game.GetDevelopmentCards()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	terrainDocuments, err := slices.Map(func(terrain *models.Terrain) (*documents.Terrain, error) {
		return terrainMapper{}.ToDocument(terrain)
	}, game.GetTerrains()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	landDocuments, err := slices.Map(func(land *models.Land) (*documents.Land, error) {
		return landMapper{}.ToDocument(land)
	}, game.GetLands()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pathDocuments, err := slices.Map(func(path *models.Path) (*documents.Path, error) {
		return pathMapper{}.ToDocument(path)
	}, game.GetPaths()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

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
	}, nil
}

func (g GameMapper) ToDomain(gameDocument *documents.Game) (*models.Game, error) {
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

	activePlayer, err := playerMapper{}.ToDomain(gameDocument.ActivePlayer)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	players, err := slices.Map(func(playerDocument *documents.Player) (*models.Player, error) {
		return playerMapper{}.ToDomain(playerDocument)
	}, gameDocument.Players...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dices, _ := slices.Map(func(diceDocument *documents.Dice) (*models.Dice, error) {
		return diceMapper{}.ToDomain(diceDocument)
	}, gameDocument.Dices...)

	achievements, err := slices.Map(func(achievementDocument *documents.Achievement) (*models.Achievement, error) {
		return achievementMapper{}.ToDomain(achievementDocument)
	}, gameDocument.Achievements...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCards, err := slices.Map(func(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
		return resourceCardMapper{}.ToDomain(resourceCardDocument)
	}, gameDocument.ResourceCards...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCards, err := slices.Map(func(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
		return developmentCardMapper{}.ToDomain(developmentCardDocument)
	}, gameDocument.DevelopmentCards...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	terrains, err := slices.Map(func(terrainDocument *documents.Terrain) (*models.Terrain, error) {
		return terrainMapper{}.ToDomain(terrainDocument)
	}, gameDocument.Terrains...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	lands, err := slices.Map(func(landDocument *documents.Land) (*models.Land, error) {
		return landMapper{}.ToDomain(landDocument)
	}, gameDocument.Lands...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	paths, err := slices.Map(func(pathDocument *documents.Path) (*models.Path, error) {
		return pathMapper{}.ToDomain(pathDocument)
	}, gameDocument.Paths...)
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
