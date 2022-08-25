package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/documents"
)

func toRoadDocument(road *models.Road) *documents.Road {
	if road == nil {
		return nil
	}

	pathDocument := toPathDocument(road.GetPath())

	return &documents.Road{
		Document: documents.Document{
			ID: road.GetID(),
		},
		Path: pathDocument,
	}
}

func toRoadDomain(roadDocument *documents.Road) (*models.Road, error) {
	if roadDocument == nil {
		return nil, nil
	}

	path, err := toPathDomain(roadDocument.Path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	road := models.NewRoadBuilder().
		SetID(roadDocument.ID).
		SetPath(path).
		Create()

	return road, nil
}
