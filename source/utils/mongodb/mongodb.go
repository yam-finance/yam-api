package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var dbUser string
var dbPass string
var dbName string
var dbURI string

type PunkIndex struct {
	ID        primitive.ObjectID `bson:"_id"`
	price     string             `bson:"price"`
	timestamp string             `bson:"timestamp"`
}

// --------- MongoDB Connection ---------

func Connect() {
	/// @dev Load .env file
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
	dbURI = os.Getenv("URI")

	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?retryWrites=true&w=majority", dbUser, dbPass, dbURI, dbName)

	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Client", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Connect", err)
	}

	// defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Ping", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

// --------- MongoDB Calls ---------

func InsertPunkIndex(val map[string]interface{}) {
	if client != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database(dbName)
		punkCollection := databaseRef.Collection("upunks")

		/// @dev Insert new values
		_, err := punkCollection.InsertOne(ctx, bson.D{
			{Key: "price", Value: val["price"]},
			{Key: "timestamp", Value: val["timestamp"]},
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetLatestPunkIndex() map[string]interface{} {
	if client != nil {
		// result := PunkIndex{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database(dbName)
		punkCollection := databaseRef.Collection("upunks")

		// @dev Find a document that meets certain criteria
		// err := punkCollection.FindOne(ctx, bson.M{}).Decode(&result)
		// if err != nil {
		// 	log.Fatal(err)
		// 	return nil
		// }

		// @dev Find last document in collection
		filterCursor, err := punkCollection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
			return nil
		}
		var resultsFiltered []bson.M
		if err = filterCursor.All(ctx, &resultsFiltered); err != nil {
			log.Fatal(err)
			return nil
		} else if len(resultsFiltered) == 0 {
			return nil
		}

		fmt.Println("Price", resultsFiltered[len(resultsFiltered)-1]["price"])
		fmt.Println("Timestamp", resultsFiltered[len(resultsFiltered)-1]["timestamp"])

		price := resultsFiltered[len(resultsFiltered)-1]["price"]
		timestamp := resultsFiltered[len(resultsFiltered)-1]["timestamp"]

		values := map[string]interface{}{"price": price, "timestamp": timestamp}

		return values
	} else {
		return nil
	}
}

func GetPunkIndexHistoryDaily() []map[string]interface{} {
	if client != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database(dbName)
		punkCollection := databaseRef.Collection("upunks")

		// @dev Find last document in collection
		filterCursor, err := punkCollection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
			return nil
		}
		var resultsFiltered []bson.M
		if err = filterCursor.All(ctx, &resultsFiltered); err != nil {
			log.Fatal(err)
			return nil
		}

		var values []map[string]interface{}
		dayCount := 0

		for _, result := range resultsFiltered {
			unixTimestamp, _ := strconv.Atoi(result["timestamp"].(string))
			timeT := time.Unix(int64(unixTimestamp), 0).UTC()
			// fmt.Printf("time.Time: %s\n", timeT)

			if strings.Contains(timeT.String(), "01:00") && dayCount < 30 {
				obj := map[string]interface{}{
					"price":     result["price"].(string),
					"timestamp": result["timestamp"].(string),
				}

				dayCount = dayCount + 1
				values = append(values, obj)
			}
		}

		return values
	} else {
		return nil
	}
}
