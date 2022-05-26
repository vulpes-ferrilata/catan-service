package mappers

import (
	"context"
	"sync"

	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/persistence/entities"
)

type PlayerMapper interface {
	ToEntity(ctx context.Context, player *models.Player, game *models.Game) (*entities.Player, error)
	ToModel(ctx context.Context, playerEntity *entities.Player) (*models.Player, error)
}

func NewPlayerMapper() PlayerMapper {
	return &playerMapper{
		m: make(map[*models.Player]*entities.Player),
	}
}

type playerMapper struct {
	m  map[*models.Player]*entities.Player
	mu sync.RWMutex
}

func (p *playerMapper) ToEntity(ctx context.Context, player *models.Player, game *models.Game) (*entities.Player, error) {
	if player == nil {
		return nil, nil
	}

	p.mu.RLock()
	playerEntity, ok := p.m[player]
	p.mu.RUnlock()
	if !ok {
		playerEntity = new(entities.Player)

		p.mu.Lock()
		p.m[player] = playerEntity
		p.mu.Unlock()

		go func(player *models.Player, done <-chan struct{}) {
			<-done
			p.mu.Lock()
			delete(p.m, player)
			p.mu.Unlock()
		}(player, ctx.Done())
	}

	playerEntity.ID = player.GetID()
	playerEntity.GameID = game.GetID()
	playerEntity.UserID = player.GetUserID()

	return playerEntity, nil
}

func (p *playerMapper) ToModel(ctx context.Context, playerEntity *entities.Player) (*models.Player, error) {
	if playerEntity == nil {
		return nil, nil
	}

	player := models.NewPlayer(
		playerEntity.ID,
		playerEntity.UserID,
	)

	p.mu.Lock()
	p.m[player] = playerEntity
	p.mu.Unlock()

	go func(player *models.Player, done <-chan struct{}) {
		<-done
		p.mu.Lock()
		delete(p.m, player)
		p.mu.Unlock()
	}(player, ctx.Done())

	return player, nil
}
