package services

import (
	"context"
	"ddaniel27/projectcore/internal/core/dto"
	"ddaniel27/projectcore/internal/core/models"
)

type ProjectsService interface {
	GetProjectByID(ctx context.Context, id int) (models.ProjectDocument, error)
	GetProjectsByOwnerID(ctx context.Context, ownerID int) ([]models.ProjectDocument, error)
	CreateProject(ctx context.Context, project *dto.CreateProjectDTO) (models.ProjectDocument, error)
	UpdateProject(ctx context.Context, project models.ProjectDocument) (models.ProjectDocument, error)
	DeleteProjectByID(ctx context.Context, id int) error
}
