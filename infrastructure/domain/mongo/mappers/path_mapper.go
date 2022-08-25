package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

func toPathDocument(path *models.Path) *documents.Path {
	if path == nil {
		return nil
	}

	pathDocument := &documents.Path{
		Document: documents.Document{
			ID: path.GetID(),
		},
		Q:        path.GetHexEdge().GetQ(),
		R:        path.GetHexEdge().GetR(),
		Location: string(path.GetHexEdge().GetLocation()),
	}

	return pathDocument
}

func toPathDomain(pathDocument *documents.Path) (*models.Path, error) {
	if pathDocument == nil {
		return nil, nil
	}

	hexEdgeLocation, err := models.NewHexEdgeLocation(pathDocument.Location)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	hexEdge := models.NewHexEdge(pathDocument.Q, pathDocument.R, hexEdgeLocation)

	path := models.NewPathBuilder().
		SetID(pathDocument.ID).
		SetHexEdge(hexEdge).
		Create()

	return path, nil
}
