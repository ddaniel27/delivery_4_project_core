package api

import (
	"context"
	"ddaniel27/projectcore/internal/infrastructure/database"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const projectCollectionName = "projects"

func (a *App) setupInfrastructure() {
	ctx := context.Background()

	db, err := newDB(ctx, "projectcore")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	collection := db.Collection(projectCollectionName)
	pr := database.NewProjectsStorage(collection)

	a.infra = &infrastructures{
		ProjectsRepository: pr,
	}
}

func newDB(ctx context.Context, dbName string) (*mongo.Database, error) {
	client, err := newClient(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}

// newClient config new client for mongo db.
func newClient(ctx context.Context) (*mongo.Client, error) {
	user := getEnvFallback("MONGO_USER", "root")
	password := getEnvFallback("MONGO_PASSWORD", "123")
	host := getEnvFallback("MONGO_HOST", "localhost")
	port := getEnvFallback("MONGO_PORT", "27017")

	check := getEnvFallback("ENVCHECK", "local")

	fmt.Println("check: ", check)

	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
	opts := options.Client().
		ApplyURI(dsn)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
