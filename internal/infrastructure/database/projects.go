package database

import (
	"context"
	"ddaniel27/projectcore/internal/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectsStorage struct {
	collection *mongo.Collection
}

func NewProjectsStorage(collection *mongo.Collection) *ProjectsStorage {
	return &ProjectsStorage{
		collection: collection,
	}
}

func (s *ProjectsStorage) GetProjectByID(ctx context.Context, id int) (models.ProjectDocument, error) {
	var project models.ProjectDocument
	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&project)

	return project, err
}

func (s *ProjectsStorage) GetProjectsByOwnerID(ctx context.Context, ownerID int) ([]models.ProjectDocument, error) {
	var projects []models.ProjectDocument
	cursor, err := s.collection.Find(ctx, bson.M{"owner.id": ownerID})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (s *ProjectsStorage) CreateProject(ctx context.Context, project models.ProjectDocument) (models.ProjectDocument, error) {
	_, err := s.collection.InsertOne(ctx, project)

	return project, err
}

func (s *ProjectsStorage) UpdateProject(ctx context.Context, project models.ProjectDocument) (models.ProjectDocument, error) {
	_, err := s.collection.ReplaceOne(ctx, bson.M{"id": project.ID}, project)

	return project, err
}

func (s *ProjectsStorage) DeleteProjectByID(ctx context.Context, id int) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"id": id})

	return err
}
