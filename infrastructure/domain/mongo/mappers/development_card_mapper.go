package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type developmentCardMapper struct{}

func (d developmentCardMapper) ToDocument(developmentCard *models.DevelopmentCard) (*documents.DevelopmentCard, error) {
	if developmentCard == nil {
		return nil, nil
	}

	return &documents.DevelopmentCard{
		Document: documents.Document{
			ID: developmentCard.GetID(),
		},
		Type:   developmentCard.GetType().String(),
		Status: string(developmentCard.GetStatus()),
	}, nil
}

func (d developmentCardMapper) ToDomain(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
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

	developmentCard := models.DevelopmentCardBuilder{}.
		SetID(developmentCardDocument.ID).
		SetType(developmentCardType).
		SetStatus(developmentCardStatus).
		Create()

	return developmentCard, nil
}
