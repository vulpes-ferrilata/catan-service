package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type terrainMapper struct{}

func (t terrainMapper) ToDocument(terrain *models.Terrain) (*documents.Terrain, error) {
	if terrain == nil {
		return nil, nil
	}

	harborDocument, err := harborMapper{}.ToDocument(terrain.GetHarbor())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	robberDocument, err := robberMapper{}.ToDocument(terrain.GetRobber())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &documents.Terrain{
		Document: documents.Document{
			ID: terrain.GetID(),
		},
		Q:      terrain.GetHex().GetQ(),
		R:      terrain.GetHex().GetR(),
		Number: terrain.GetNumber(),
		Type:   string(terrain.GetType()),
		Harbor: harborDocument,
		Robber: robberDocument,
	}, nil
}

func (t terrainMapper) ToDomain(terrainDocument *documents.Terrain) (*models.Terrain, error) {
	if terrainDocument == nil {
		return nil, nil
	}

	hex := models.NewHex(terrainDocument.Q, terrainDocument.R)

	terrainType, err := models.NewTerrainType(terrainDocument.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	harbor, err := harborMapper{}.ToDomain(terrainDocument.Harbor)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	robber, err := robberMapper{}.ToDomain(terrainDocument.Robber)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	terrain := models.TerrainBuilder{}.
		SetID(terrainDocument.ID).
		SetHex(hex).
		SetNumber(terrainDocument.Number).
		SetType(terrainType).
		SetHarbor(harbor).
		SetRobber(robber).
		Create()

	return terrain, nil
}
