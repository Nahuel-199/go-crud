package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	// Carga las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtiene el URI de la base de datos desde las variables de entorno
	mongoURI := os.Getenv("MONGO_URI")

	// Crea una instancia del cliente de MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verifica la conexión
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	return client, nil
}

func GetCollection(collectionName string) *mongo.Collection {
	client, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("goback").Collection(collectionName)
	return collection
}
