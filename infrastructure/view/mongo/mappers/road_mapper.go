package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type roadMapper struct{}

func (r roadMapper) ToView(roadDocument *documents.Road) (*models.Road, error) {
	if roadDocument == nil {
		return nil, nil
	}

	path, err := pathMapper{}.ToView(roadDocument.Path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.Road{
		ID:   roadDocument.ID,
		Path: path,
	}, nil
}
