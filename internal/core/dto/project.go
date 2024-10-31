package dto

import "ddaniel27/projectcore/internal/core/models"

type CreateProjectDTO struct {
	Name        string              `json:"name" binding:"required"`
	Owner       models.Owner        `json:"owner" binding:"required"`
	Groups      []models.Group      `json:"groups" binding:"required"`
	Nominations []models.Nomination `json:"nominations" binding:"required"`
	IsPublic    bool                `json:"isPublic" binding:"required"`
}
