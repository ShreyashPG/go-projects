package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB     *mongo.Client
	once   sync.Once
	dbName = "employees" // Define database name here
)

// ConnectToDB initializes MongoDB connection
func ConnectToDB() error {
	once.Do(func() { // Ensures DB is initialized only once
		// Load .env file if not already loaded
		_ = godotenv.Load()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		uri := os.Getenv("MONGO_URI")
		if uri == "" {
			log.Fatal("MONGO_URI not set in environment")
		}

		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal("Failed to ping MongoDB:", err)
		}

		DB = client
		log.Println("✅ Connected to MongoDB")
	})

	return nil
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB() {
	if DB != nil {
		err := DB.Disconnect(context.Background())
		if err != nil {
			log.Println("Error disconnecting MongoDB:", err)
		} else {
			log.Println("🛑 Disconnected from MongoDB")
		}
	}
}

// GetCollection returns a reference to a MongoDB collection
func GetCollection(name string) *mongo.Collection {
	if DB == nil {
		log.Fatal("MongoDB not initialized. Call ConnectToDB() first.")
	}
	return DB.Database(dbName).Collection(name)
}
