package services

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/domain/mappers"
	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type GameService interface {
	NewGame(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*models.Game, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Game, error)
	Save(ctx context.Context, game *models.Game) error
}

func NewGameService(gameRepository repositories.GameRepository,
	gameMapper mappers.GameMapper,
	playerService PlayerService,
	diceService DiceService,
	achievementService AchievementService) GameService {
	return &gameService{
		gameRepository:     gameRepository,
		gameMapper:         gameMapper,
		playerService:      playerService,
		diceService:        diceService,
		achievementService: achievementService,
	}
}

type gameService struct {
	gameRepository     repositories.GameRepository
	gameMapper         mappers.GameMapper
	playerService      PlayerService
	diceService        DiceService
	achievementService AchievementService
}

func (g gameService) NewGame(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*models.Game, error) {
	player, err := g.playerService.NewPlayer(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dices, err := g.diceService.NewDices(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievements, err := g.achievementService.NewAchievements(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	game := models.NewGame(id, models.Waiting, player.GetID(), 1, []*models.Player{player}, dices, achievements)

	return game, nil
}

func (g gameService) GetByID(ctx context.Context, id uuid.UUID) (*models.Game, error) {
	gameEntity, err := g.gameRepository.GetByID(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	players, err := g.playerService.FindByGameID(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dices, err := g.diceService.FindByGameID(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievements, err := g.achievementService.FindByGameIDByPlayerID(ctx, id, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	game, err := g.gameMapper.ToModel(ctx, gameEntity, players, dices, achievements)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return game, nil
}

func (g gameService) Save(ctx context.Context, game *models.Game) error {
	gameEntity, err := g.gameMapper.ToEntity(ctx, game)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := g.gameRepository.Save(ctx, gameEntity); err != nil {
		return errors.WithStack(err)
	}

	for _, player := range game.GetPlayers() {
		if err := g.playerService.Save(ctx, player, game); err != nil {
			return errors.WithStack(err)
		}
	}

	for _, dice := range game.GetDices() {
		if err := g.diceService.Save(ctx, dice, game); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
