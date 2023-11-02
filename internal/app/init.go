package app

import (
	"context"
	"flag"
	"time"

	"github.com/math-88/medic-online/internal/entity"
	"github.com/math-88/medic-online/storage"

	"github.com/google/uuid"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint    = flag.String("grpc-server-endpoint", ":8080", "gRPC server endpoint")
	restServerEndpoint    = flag.String("rest-server-endpoint", ":8081", "REST server endpoint")
	graphQLServerEndpoint = flag.String("graphql-server-endpoint", ":8082", "GraphQL server endpoint")
)

func seedUsers(userStore storage.Storage) error {
	err := createUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user")
}

func createUser(userStore storage.Storage, username, password, role string) error {

	user, err := entity.NewUser(uuid.New(), username, password, role)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return userStore.Save(ctx, user)
}

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)
