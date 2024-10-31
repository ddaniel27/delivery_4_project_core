package projects

import (
	"context"
	"ddaniel27/projectcore/internal/core/models"
)

func (ps *ProjectService) UpdateProject(_ context.Context, _ models.ProjectDocument) (models.ProjectDocument, error) {
	return models.ProjectDocument{}, nil
}
