package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type landMapper struct{}

func (l landMapper) ToDocument(land *models.Land) (*documents.Land, error) {
	if land == nil {
		return nil, nil
	}

	landDocument := &documents.Land{
		Document: documents.Document{
			ID: land.GetID(),
		},
		Q:        land.GetHexCorner().GetQ(),
		R:        land.GetHexCorner().GetR(),
		Location: land.GetHexCorner().GetLocation().String(),
	}

	return landDocument, nil
}

func (l landMapper) ToDomain(landDocument *documents.Land) (*models.Land, error) {
	if landDocument == nil {
		return nil, nil
	}

	hexCornerLocation, err := models.NewHexCornerLocation(landDocument.Location)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	hexCorner := models.NewHexCorner(landDocument.Q, landDocument.R, hexCornerLocation)

	land := models.LandBuilder{}.
		SetID(landDocument.ID).
		SetHexCorner(hexCorner).
		Create()

	return land, nil
}
