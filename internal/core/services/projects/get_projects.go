package projects

import (
	"context"
	"ddaniel27/projectcore/internal/core/models"
)

func (ps *ProjectService) GetProjectByID(ctx context.Context, id int) (models.ProjectDocument, error) {
	p, err := ps.ProjectRepository.GetProjectByID(ctx, id)
	if err != nil {
		return models.ProjectDocument{}, err
	}

	return p, nil
}

func (ps *ProjectService) GetProjectsByOwnerID(ctx context.Context, ownerID int) ([]models.ProjectDocument, error) {
	projects, err := ps.ProjectRepository.GetProjectsByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
