package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
)

func toPlayerDocument(player *models.Player) *documents.Player {
	if player == nil {
		return nil
	}

	achievementDocuments, _ := slices.Map(func(achievement *models.Achievement) (*documents.Achievement, error) {
		return toAchievementDocument(achievement), nil
	}, player.GetAchievements())

	resourceCardDocuments, _ := slices.Map(func(resourceCard *models.ResourceCard) (*documents.ResourceCard, error) {
		return toResourceCardDocument(resourceCard), nil
	}, player.GetResourceCards())

	developmentCardDocuments, _ := slices.Map(func(developmentCard *models.DevelopmentCard) (*documents.DevelopmentCard, error) {
		return toDevelopmentCardDocument(developmentCard), nil
	}, player.GetDevelopmentCards())

	constructionDocuments, _ := slices.Map(func(construction *models.Construction) (*documents.Construction, error) {
		return toConstructionDocument(construction), nil
	}, player.GetConstructions())

	roadDocuments, _ := slices.Map(func(road *models.Road) (*documents.Road, error) {
		return toRoadDocument(road), nil
	}, player.GetRoads())

	return &documents.Player{
		Document: documents.Document{
			ID: player.GetID(),
		},
		UserID:           player.GetUserID(),
		TurnOrder:        player.GetTurnOrder(),
		Color:            string(player.GetColor()),
		IsActive:         player.IsActive(),
		IsOffered:        player.IsOffered(),
		Achievements:     achievementDocuments,
		ResourceCards:    resourceCardDocuments,
		DevelopmentCards: developmentCardDocuments,
		Constructions:    constructionDocuments,
		Roads:            roadDocuments,
	}
}

func toPlayerDomain(playerDocument *documents.Player) (*models.Player, error) {
	if playerDocument == nil {
		return nil, nil
	}

	color, err := models.NewPlayerColor(playerDocument.Color)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievements, err := slices.Map(func(achievementDocument *documents.Achievement) (*models.Achievement, error) {
		return toAchievementDomain(achievementDocument)
	}, playerDocument.Achievements)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCards, err := slices.Map(func(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
		return toResourceCardDomain(resourceCardDocument)
	}, playerDocument.ResourceCards)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCards, err := slices.Map(func(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
		return toDevelopmentCardDomain(developmentCardDocument)
	}, playerDocument.DevelopmentCards)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	constructions, err := slices.Map(func(constructionDocument *documents.Construction) (*models.Construction, error) {
		return toConstructionDomain(constructionDocument)
	}, playerDocument.Constructions)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	roads, err := slices.Map(func(roadDocument *documents.Road) (*models.Road, error) {
		return toRoadDomain(roadDocument)
	}, playerDocument.Roads)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	player := models.NewPlayerBuilder().
		SetID(playerDocument.ID).
		SetUserID(playerDocument.UserID).
		SetColor(color).
		SetTurnOrder(playerDocument.TurnOrder).
		SetIsActive(playerDocument.IsActive).
		SetIsOffered(playerDocument.IsOffered).
		SetAchievements(achievements...).
		SetResourceCards(resourceCards...).
		SetDevelopmentCards(developmentCards...).
		SetConstructions(constructions...).
		SetRoads(roads...).
		Create()

	return player, nil
}
