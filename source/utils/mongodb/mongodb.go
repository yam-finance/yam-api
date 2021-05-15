package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

type AprYam struct {
	ID    primitive.ObjectID `bson:"_id"`
	Value float64            `bson:"apryam"`
}
type AprDegenerative struct {
	ID    primitive.ObjectID `bson:"_id"`
	Value map[string]float64 `bson:"aprdegenerative"`
}

func Connect() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbURI := os.Getenv("URI")

	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?retryWrites=true&w=majority", dbUser, dbPass, dbURI, dbName)
	fmt.Println(uri)
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Client", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Connect", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Ping", err)
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

}
func InsertAprYam(val float64) {
	if client != nil {
		result := AprYam{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database("ugasmedian")
		aprYamCollection := databaseRef.Collection("apryam")
		update := bson.M{
			"$set": bson.M{"apryam": val},
		}
		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}
		er := aprYamCollection.FindOneAndUpdate(ctx, bson.M{}, update, &opt).Decode(&result)
		if er != nil {
			fmt.Println(er)
			return
		}
	}
}
func InsertAprDegenerative(val map[string]float64) {
	if client != nil {
		result := AprDegenerative{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database("ugasmedian")
		aprYamDegenerativeCollection := databaseRef.Collection("aprdegenerative")
		update := bson.M{
			"$set": bson.M{"aprdegenerative": val},
		}
		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}
		er := aprYamDegenerativeCollection.FindOneAndUpdate(ctx, bson.M{}, update, &opt).Decode(&result)
		if er != nil {
			fmt.Println(er)
			return
		}
	}
}
func GetAprYam() float64 {
	if client != nil {
		result := AprYam{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database("ugasmedian")
		aprYamCollection := databaseRef.Collection("apryam")
		er := aprYamCollection.FindOne(ctx, bson.M{}).Decode(&result)
		if er != nil {
			fmt.Println(er)
			return 0
		}
		fmt.Println(result.Value)
		return result.Value
	}
	return 0
}
func GetAprDegenerative() map[string]float64 {
	if client != nil {
		result := AprDegenerative{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database("ugasmedian")
		aprYamDegenerativeCollection := databaseRef.Collection("aprdegenerative")
		er := aprYamDegenerativeCollection.FindOne(ctx, bson.M{}).Decode(&result)
		if er != nil {
			fmt.Println(er)
			return map[string]float64{"MAR21": 0, "JUN21": 0}
		}
		fmt.Println(result.Value)
		return result.Value
	}
	return map[string]float64{"MAR21": 0, "JUN21": 0}
}
