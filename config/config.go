package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDB initializes the MongoDB client and connects to the database.
func ConnectDB() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://gouser:gopassword@cluster0.zichifw.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    DB = client.Database("todoapp")
    log.Println("Connected to MongoDB!")
}
