package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

type roadMapper struct{}

func (r roadMapper) ToDocument(road *models.Road) (*documents.Road, error) {
	if road == nil {
		return nil, nil
	}

	pathDocument, err := pathMapper{}.ToDocument(road.GetPath())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &documents.Road{
		Document: documents.Document{
			ID: road.GetID(),
		},
		Path: pathDocument,
	}, nil
}

func (r roadMapper) ToDomain(roadDocument *documents.Road) (*models.Road, error) {
	if roadDocument == nil {
		return nil, nil
	}

	path, err := pathMapper{}.ToDomain(roadDocument.Path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	road := models.RoadBuilder{}.
		SetID(roadDocument.ID).
		SetPath(path).
		Create()

	return road, nil
}
