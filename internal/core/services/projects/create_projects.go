package projects

import (
	"context"
	"ddaniel27/projectcore/internal/core/dto"
	"ddaniel27/projectcore/internal/core/models"
	"time"
)

func (ps *ProjectService) CreateProject(ctx context.Context, p *dto.CreateProjectDTO) (models.ProjectDocument, error) {
	project := models.ProjectDocument{
		Name:        p.Name,
		Groups:      p.Groups,
		Nominations: p.Nominations,
		Owner:       p.Owner,
		Metadata: models.Metadata{
			IsPublic:  p.IsPublic,
			CreatedAt: time.Now().UTC().Unix(),
			UpdatedAt: time.Now().UTC().Unix(),
		},
	}

	project, err := ps.ProjectRepository.CreateProject(ctx, project)
	if err != nil {
		return models.ProjectDocument{}, err
	}

	return project, nil
}
