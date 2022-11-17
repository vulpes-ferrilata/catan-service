package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type pathMapper struct{}

func (p pathMapper) ToDocument(path *models.Path) (*documents.Path, error) {
	if path == nil {
		return nil, nil
	}

	pathDocument := &documents.Path{
		Document: documents.Document{
			ID: path.GetID(),
		},
		Q:        path.GetHexEdge().GetQ(),
		R:        path.GetHexEdge().GetR(),
		Location: path.GetHexEdge().GetLocation().String(),
	}

	return pathDocument, nil
}

func (p pathMapper) ToDomain(pathDocument *documents.Path) (*models.Path, error) {
	if pathDocument == nil {
		return nil, nil
	}

	hexEdgeLocation, err := models.NewHexEdgeLocation(pathDocument.Location)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	hexEdge := models.NewHexEdge(pathDocument.Q, pathDocument.R, hexEdgeLocation)

	path := models.PathBuilder{}.
		SetID(pathDocument.ID).
		SetHexEdge(hexEdge).
		Create()

	return path, nil
}
