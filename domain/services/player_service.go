package services

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/domain/mappers"
	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type PlayerService interface {
	NewPlayer(ctx context.Context, userID uuid.UUID) (*models.Player, error)
	FindByGameID(ctx context.Context, gameID uuid.UUID) ([]*models.Player, error)
	Save(ctx context.Context, player *models.Player, game *models.Game) error
}

func NewPlayerService(playerRepository repositories.PlayerRepository,
	playerMapper mappers.PlayerMapper) PlayerService {
	return &playerService{
		playerRepository: playerRepository,
		playerMapper:     playerMapper,
	}
}

type playerService struct {
	playerRepository repositories.PlayerRepository
	playerMapper     mappers.PlayerMapper
}

func (p playerService) NewPlayer(ctx context.Context, userID uuid.UUID) (*models.Player, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	player := models.NewPlayer(id, userID)

	return player, nil
}

func (p playerService) FindByGameID(ctx context.Context, gameID uuid.UUID) ([]*models.Player, error) {
	players := make([]*models.Player, 0)

	playerEntities, err := p.playerRepository.FindByGameID(ctx, gameID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, playerEntity := range playerEntities {
		player, err := p.playerMapper.ToModel(ctx, playerEntity)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		players = append(players, player)
	}

	return players, nil
}

func (p playerService) Save(ctx context.Context, player *models.Player, game *models.Game) error {
	playerEntity, err := p.playerMapper.ToEntity(ctx, player, game)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := p.playerRepository.Save(ctx, playerEntity); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
