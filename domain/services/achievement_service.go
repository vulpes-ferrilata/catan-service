package services

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/domain/mappers"
	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AchievementService interface {
	NewAchievements(ctx context.Context) ([]*models.Achievement, error)
	FindByGameIDByPlayerID(ctx context.Context, gameID uuid.UUID, playerID *uuid.UUID) ([]*models.Achievement, error)
	Save(ctx context.Context, achievement *models.Achievement, game *models.Game, player *models.Player) error
}

func NewAchievementService(achievementRepository repositories.AchievementRepository,
	achievementMapper mappers.AchievementMapper) AchievementService {
	return &achievementService{
		achievementRepository: achievementRepository,
		achievementMapper:     achievementMapper,
	}
}

type achievementService struct {
	achievementRepository repositories.AchievementRepository
	achievementMapper     mappers.AchievementMapper
}

func (a achievementService) NewAchievement(ctx context.Context, achievementType models.AchievementType) (*models.Achievement, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievement := models.NewAchievement(id, achievementType)

	return achievement, nil
}

func (a achievementService) NewAchievements(ctx context.Context) ([]*models.Achievement, error) {
	achievements := make([]*models.Achievement, 0)

	longestRoadAchievement, err := a.NewAchievement(ctx, models.LongestRoad)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	achievements = append(achievements, longestRoadAchievement)

	largestArmyAchievement, err := a.NewAchievement(ctx, models.LargestArmy)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	achievements = append(achievements, largestArmyAchievement)

	return achievements, nil
}

func (a achievementService) FindByGameIDByPlayerID(ctx context.Context, gameID uuid.UUID, playerID *uuid.UUID) ([]*models.Achievement, error) {
	achievements := make([]*models.Achievement, 0)

	achievementEntities, err := a.achievementRepository.FindByGameIDByPlayerID(ctx, gameID, playerID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, achievementEntity := range achievementEntities {
		achievement, err := a.achievementMapper.ToModel(ctx, achievementEntity)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		achievements = append(achievements, achievement)
	}

	return achievements, nil
}

func (a achievementService) Save(ctx context.Context, achievement *models.Achievement, game *models.Game, player *models.Player) error {
	achievementEntity, err := a.achievementMapper.ToEntity(ctx, achievement, game, player)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := a.achievementRepository.Save(ctx, achievementEntity); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
