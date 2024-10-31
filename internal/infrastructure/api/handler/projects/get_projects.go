package handler

import "github.com/gin-gonic/gin"

type pathID struct {
	ID int `uri:"id" binding:"required"`
}

func (h *ProjectHandler) GetProjectByID(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "GetProjectByID")
	defer span.End()

	var p pathID
	if err := c.ShouldBindUri(&p); err != nil {
		h.logger.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	project, err := h.ProjectsService.GetProjectByID(ctx, p.ID)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"project": project})
}

func (h *ProjectHandler) GetProjectsByOwnerID(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "GetProjectsByOwnerID")
	defer span.End()

	var p pathID
	if err := c.ShouldBindUri(&p); err != nil {
		h.logger.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	project, err := h.ProjectsService.GetProjectsByOwnerID(ctx, p.ID)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"project": project})
}
