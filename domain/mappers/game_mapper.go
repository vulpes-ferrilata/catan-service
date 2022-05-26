package mappers

import (
	"context"
	"sync"

	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/persistence/entities"
	"github.com/pkg/errors"
)

type GameMapper interface {
	ToEntity(ctx context.Context, game *models.Game) (*entities.Game, error)
	ToModel(ctx context.Context, gameEntity *entities.Game, players []*models.Player, dices []*models.Dice, achievements []*models.Achievement) (*models.Game, error)
}

func NewGameMapper() GameMapper {
	return &gameMapper{
		m: make(map[*models.Game]*entities.Game),
	}
}

type gameMapper struct {
	m  map[*models.Game]*entities.Game
	mu sync.RWMutex
}

func (g *gameMapper) ToEntity(ctx context.Context, game *models.Game) (*entities.Game, error) {
	if game == nil {
		return nil, nil
	}

	g.mu.RLock()
	gameEntity, ok := g.m[game]
	g.mu.RUnlock()
	if !ok {
		gameEntity = new(entities.Game)

		g.mu.Lock()
		g.m[game] = gameEntity
		g.mu.Unlock()

		go func(game *models.Game, done <-chan struct{}) {
			<-done
			g.mu.Lock()
			delete(g.m, game)
			g.mu.Unlock()
		}(game, ctx.Done())
	}

	gameEntity.ID = game.GetID()
	gameEntity.Status = string(game.GetStatus())
	gameEntity.ActivePlayerID = game.GetActivePlayerID()
	gameEntity.Turn = game.GetTurn()

	return gameEntity, nil
}

func (g *gameMapper) ToModel(ctx context.Context, gameEntity *entities.Game, players []*models.Player, dices []*models.Dice, achievements []*models.Achievement) (*models.Game, error) {
	if gameEntity == nil {
		return nil, nil
	}

	gameStatus, err := models.NewGameStatus(gameEntity.Status)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	game := models.NewGame(
		gameEntity.ID,
		gameStatus,
		gameEntity.ActivePlayerID,
		gameEntity.Turn,
		players,
		dices,
		achievements,
	)

	g.mu.Lock()
	g.m[game] = gameEntity
	g.mu.Unlock()

	go func(game *models.Game, done <-chan struct{}) {
		<-done
		g.mu.Lock()
		delete(g.m, game)
		g.mu.Unlock()
	}(game, ctx.Done())

	return game, nil
}
