package projects

import "ddaniel27/projectcore/internal/core/ports/repositories"

type ProjectService struct {
	ProjectRepository repositories.ProjectsRepository
}

func NewProjectService(pr repositories.ProjectsRepository) *ProjectService {
	return &ProjectService{
		ProjectRepository: pr,
	}
}
