package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

func toRobberDocument(robber *models.Robber) *documents.Robber {
	if robber == nil {
		return nil
	}

	return &documents.Robber{
		Document: documents.Document{
			ID: robber.GetID(),
		},
		TerrainID: robber.GetTerrainID(),
		IsMoving:  robber.IsMoving(),
	}
}

func toRobberDomain(robberDocument *documents.Robber) (*models.Robber, error) {
	if robberDocument == nil {
		return nil, nil
	}

	robber := models.NewRobberBuilder().
		SetID(robberDocument.ID).
		SetTerrainID(robberDocument.TerrainID).
		SetIsMoving(robberDocument.IsMoving).
		Create()

	return robber, nil
}
