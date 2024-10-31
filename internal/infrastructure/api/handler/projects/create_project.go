package handler

import (
	"ddaniel27/projectcore/internal/core/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "CreateProject")
	defer span.End()

	var body dto.CreateProjectDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		h.logger.Error(fmt.Sprintf("Failed to bind JSON: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.ProjectsService.CreateProject(ctx, &body); err != nil {
		h.logger.Error(fmt.Sprintf("Failed to create project: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"project_created": false, "error": err.Error()})
		return
	}

	h.projectCounter.Add(ctx, 1)
	c.JSON(http.StatusCreated, gin.H{"project_created": true})
}
