package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
)

type playerMapper struct{}

func (p playerMapper) ToDocument(player *models.Player) (*documents.Player, error) {
	if player == nil {
		return nil, nil
	}

	achievementDocuments, err := slices.Map(func(achievement *models.Achievement) (*documents.Achievement, error) {
		return achievementMapper{}.ToDocument(achievement)
	}, player.GetAchievements())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCardDocuments, err := slices.Map(func(resourceCard *models.ResourceCard) (*documents.ResourceCard, error) {
		return resourceCardMapper{}.ToDocument(resourceCard)
	}, player.GetResourceCards())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCardDocuments, err := slices.Map(func(developmentCard *models.DevelopmentCard) (*documents.DevelopmentCard, error) {
		return developmentCardMapper{}.ToDocument(developmentCard)
	}, player.GetDevelopmentCards())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	constructionDocuments, err := slices.Map(func(construction *models.Construction) (*documents.Construction, error) {
		return constructionMapper{}.ToDocument(construction)
	}, player.GetConstructions())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	roadDocuments, err := slices.Map(func(road *models.Road) (*documents.Road, error) {
		return roadMapper{}.ToDocument(road)
	}, player.GetRoads())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &documents.Player{
		Document: documents.Document{
			ID: player.GetID(),
		},
		UserID:             player.GetUserID(),
		TurnOrder:          player.GetTurnOrder(),
		Color:              player.GetColor().String(),
		ReceivedOffer:      player.IsReceivedOffer(),
		DiscardedResources: player.IsDiscardedResources(),
		Score:              player.GetScore(),
		Achievements:       achievementDocuments,
		ResourceCards:      resourceCardDocuments,
		DevelopmentCards:   developmentCardDocuments,
		Constructions:      constructionDocuments,
		Roads:              roadDocuments,
	}, nil
}

func (p playerMapper) ToDomain(playerDocument *documents.Player) (*models.Player, error) {
	if playerDocument == nil {
		return nil, nil
	}

	color, err := models.NewPlayerColor(playerDocument.Color)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	achievements, err := slices.Map(func(achievementDocument *documents.Achievement) (*models.Achievement, error) {
		return achievementMapper{}.ToDomain(achievementDocument)
	}, playerDocument.Achievements)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCards, err := slices.Map(func(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
		return resourceCardMapper{}.ToDomain(resourceCardDocument)
	}, playerDocument.ResourceCards)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCards, err := slices.Map(func(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
		return developmentCardMapper{}.ToDomain(developmentCardDocument)
	}, playerDocument.DevelopmentCards)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	constructions, err := slices.Map(func(constructionDocument *documents.Construction) (*models.Construction, error) {
		return constructionMapper{}.ToDomain(constructionDocument)
	}, playerDocument.Constructions)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	roads, err := slices.Map(func(roadDocument *documents.Road) (*models.Road, error) {
		return roadMapper{}.ToDomain(roadDocument)
	}, playerDocument.Roads)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	player := models.PlayerBuilder{}.
		SetID(playerDocument.ID).
		SetUserID(playerDocument.UserID).
		SetColor(color).
		SetTurnOrder(playerDocument.TurnOrder).
		SetReceivedOffer(playerDocument.ReceivedOffer).
		SetDiscardedResources(playerDocument.DiscardedResources).
		SetScore(playerDocument.Score).
		SetAchievements(achievements...).
		SetResourceCards(resourceCards...).
		SetDevelopmentCards(developmentCards...).
		SetConstructions(constructions...).
		SetRoads(roads...).
		Create()

	return player, nil
}
