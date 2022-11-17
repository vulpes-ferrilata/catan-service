package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type resourceCardMapper struct{}

func (r resourceCardMapper) ToDocument(resourceCard *models.ResourceCard) (*documents.ResourceCard, error) {
	if resourceCard == nil {
		return nil, nil
	}

	return &documents.ResourceCard{
		Document: documents.Document{
			ID: resourceCard.GetID(),
		},
		Type:     resourceCard.GetType().String(),
		Offering: resourceCard.IsOffering(),
	}, nil
}

func (r resourceCardMapper) ToDomain(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
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
		SetOffering(resourceCardDocument.Offering).
		Create()

	return resourceCard, nil
}
