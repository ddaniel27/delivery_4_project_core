package api

import (
	"context"
	"ddaniel27/projectcore/internal/core/services/projects"
	projecthandler "ddaniel27/projectcore/internal/infrastructure/api/handler/projects"
	"ddaniel27/projectcore/internal/infrastructure/observability"
	"log"
)

func (a *App) setupDependencies() {
	ps := projects.NewProjectService(a.infra.ProjectsRepository)

	ph := projecthandler.NewProjectHandler(ps)

	otelShutdown, err := observability.SetupOtelSDK(context.Background())
	if err != nil {
		log.Fatalf("Failed to setup observability: %v", err)
	}

	a.deps = &dependencies{
		ProjectsHandler: *ph,
		OtelShutdown:    otelShutdown,
	}
}
