package repositories

import (
	"context"
	"ddaniel27/projectcore/internal/core/models"
)

type ProjectsRepository interface {
	GetProjectByID(ctx context.Context, id int) (models.ProjectDocument, error)
	GetProjectsByOwnerID(ctx context.Context, ownerID int) ([]models.ProjectDocument, error)
	CreateProject(ctx context.Context, project models.ProjectDocument) (models.ProjectDocument, error)
	UpdateProject(ctx context.Context, project models.ProjectDocument) (models.ProjectDocument, error)
	DeleteProjectByID(ctx context.Context, id int) error
}
