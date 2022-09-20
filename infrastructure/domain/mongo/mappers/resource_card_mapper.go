package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

func toResourceCardDocument(resourceCard *models.ResourceCard) *documents.ResourceCard {
	if resourceCard == nil {
		return nil
	}

	return &documents.ResourceCard{
		Document: documents.Document{
			ID: resourceCard.GetID(),
		},
		Type:       resourceCard.GetType().String(),
		IsSelected: resourceCard.IsSelected(),
	}
}

func toResourceCardDomain(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
	if resourceCardDocument == nil {
		return nil, nil
	}

	resourceCardType, err := models.NewResourceCardType(resourceCardDocument.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCard := models.ResourceCardBuilder{}.
		SetID(resourceCardDocument.ID).
		SetType(resourceCardType).
		SetIsSelected(resourceCardDocument.IsSelected).
		Create()

	return resourceCard, nil
}
