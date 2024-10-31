package handler

import (
	"ddaniel27/projectcore/internal/core/ports/services"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const name = "project"

type (
	ProjectHandler struct {
		ProjectsService services.ProjectsService
		observability
	}

	observability struct {
		tracer         trace.Tracer
		meter          metric.Meter
		logger         slog.Logger
		projectCounter metric.Int64Counter
	}
)

func NewProjectHandler(s services.ProjectsService) *ProjectHandler {
	var err error

	observability := observability{
		tracer: otel.Tracer(name),
		meter:  otel.Meter(name),
		logger: *otelslog.NewLogger(name),
	}

	observability.projectCounter, err = observability.
		meter.
		Int64Counter(
			"project.counter",
			metric.WithDescription("Number of projects created"),
			metric.WithUnit("{number}"),
		)
	if err != nil {
		observability.logger.Error(fmt.Sprintf("Failed to create counter: %s", err.Error()))
	}

	return &ProjectHandler{
		ProjectsService: s,
		observability:   observability,
	}
}
