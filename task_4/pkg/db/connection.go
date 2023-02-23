package db

import (
	"context"
	"log"
	"time"

	"example.com/machine_test/task_4/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() error {
	// serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	//SetServerAPIOptions(serverAPIOptions)

	credential := options.Credential{
		Username: "admin",
		Password: "password",
	}

	clientOptions.SetAuth(credential)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	usecase.NewClientInit(client)

	log.Println("Connected to mongo db server!")
	return nil
}
