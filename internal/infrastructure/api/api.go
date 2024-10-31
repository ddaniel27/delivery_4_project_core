package api

import (
	"context"
	"ddaniel27/projectcore/internal/core/ports/repositories"
	projecthandler "ddaniel27/projectcore/internal/infrastructure/api/handler/projects"

	"github.com/gin-gonic/gin"
)

type dependencies struct {
	ProjectsHandler projecthandler.ProjectHandler
	OtelShutdown    func(context.Context) error
}

type infrastructures struct {
	ProjectsRepository repositories.ProjectsRepository
}

type App struct {
	Server *gin.Engine
	deps   *dependencies
	infra  *infrastructures
}

func NewApp() *App {
	a := &App{}
	a.setupInfrastructure()
	a.setupDependencies()
	a.setupServer()

	return a
}
