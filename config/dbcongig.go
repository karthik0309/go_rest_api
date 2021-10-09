package config

import(
	"log"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI="mongodb+srv://karthik:Kartz1234@cluster0.7tjcx.mongodb.net/instagram?retryWrites=true&w=majority"
var database = "go_rest_api"

func GetCollections(collection string) *mongo.Collection{
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

	return client.Database(database).Collection(collection)
}