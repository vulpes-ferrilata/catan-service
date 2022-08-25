package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

func toDevelopmentCardDocument(developmentCard *models.DevelopmentCard) *documents.DevelopmentCard {
	if developmentCard == nil {
		return nil
	}

	return &documents.DevelopmentCard{
		Document: documents.Document{
			ID: developmentCard.GetID(),
		},
		Type:   string(developmentCard.GetType()),
		Status: string(developmentCard.GetStatus()),
	}
}

func toDevelopmentCardDomain(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
	if developmentCardDocument == nil {
		return nil, nil
	}

	developmentCardType, err := models.NewDevelopmentCardType(developmentCardDocument.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCardStatus, err := models.NewDevelopmentCardStatus(developmentCardDocument.Status)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCard := models.NewDevelopmentCardBuilder().
		SetID(developmentCardDocument.ID).
		SetType(developmentCardType).
		SetStatus(developmentCardStatus).
		Create()

	return developmentCard, nil
}
