package mappers

import (
	"context"
	"sync"

	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/persistence/entities"
	"github.com/pkg/errors"
)

type AchievementMapper interface {
	ToEntity(ctx context.Context, achievement *models.Achievement, game *models.Game, player *models.Player) (*entities.Achievement, error)
	ToModel(ctx context.Context, achievementEntity *entities.Achievement) (*models.Achievement, error)
}

func NewAchievementMapper() AchievementMapper {
	return &achievementMapper{
		m: make(map[*models.Achievement]*entities.Achievement),
	}
}

type achievementMapper struct {
	m  map[*models.Achievement]*entities.Achievement
	mu sync.RWMutex
}

func (p *achievementMapper) ToEntity(ctx context.Context, achievement *models.Achievement, game *models.Game, player *models.Player) (*entities.Achievement, error) {
	if achievement == nil {
		return nil, nil
	}

	p.mu.RLock()
	achievementEntity, ok := p.m[achievement]
	p.mu.RUnlock()
	if !ok {
		achievementEntity = new(entities.Achievement)

		p.mu.Lock()
		p.m[achievement] = achievementEntity
		p.mu.Unlock()

		go func(achievement *models.Achievement, done <-chan struct{}) {
			<-done
			p.mu.Lock()
			delete(p.m, achievement)
			p.mu.Unlock()
		}(achievement, ctx.Done())
	}

	achievementEntity.ID = achievement.GetID()
	achievementEntity.GameID = game.GetID()
	if player != nil {
		*achievementEntity.PlayerID = player.GetID()
	} else {
		achievementEntity.PlayerID = nil
	}
	achievementEntity.Type = string(achievement.GetType())

	return achievementEntity, nil
}

func (p *achievementMapper) ToModel(ctx context.Context, achievementEntity *entities.Achievement) (*models.Achievement, error) {
	if achievementEntity == nil {
		return nil, nil
	}

	achievementType, err := models.NewAchievementType(achievementEntity.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievement := models.NewAchievement(
		achievementEntity.ID,
		achievementType,
	)

	p.mu.Lock()
	p.m[achievement] = achievementEntity
	p.mu.Unlock()

	go func(achievement *models.Achievement, done <-chan struct{}) {
		<-done
		p.mu.Lock()
		delete(p.m, achievement)
		p.mu.Unlock()
	}(achievement, ctx.Done())

	return achievement, nil
}
