package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

func toTerrainDocument(terrain *models.Terrain) *documents.Terrain {
	if terrain == nil {
		return nil
	}

	return &documents.Terrain{
		Document: documents.Document{
			ID: terrain.GetID(),
		},
		Q:      terrain.GetHex().GetQ(),
		R:      terrain.GetHex().GetR(),
		Number: terrain.GetNumber(),
		Type:   string(terrain.GetType()),
	}
}

func toTerrainDomain(terrainDocument *documents.Terrain) (*models.Terrain, error) {
	if terrainDocument == nil {
		return nil, nil
	}

	hex := models.NewHex(terrainDocument.Q, terrainDocument.R)

	terrainType, err := models.NewTerrainType(terrainDocument.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	terrain := models.NewTerrainBuilder().
		SetID(terrainDocument.ID).
		SetHex(hex).
		SetNumber(terrainDocument.Number).
		SetType(terrainType).
		Create()

	return terrain, nil
}
