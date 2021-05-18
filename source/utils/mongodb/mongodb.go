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

func Connect() {
	/// @dev Load .env file
	if _, err := os.Stat(".env"); err == nil || os.IsExist(err) {
		envErr := godotenv.Load(".env")
		if envErr != nil {
			log.Fatalf("Error loading .env file")
		}
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
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Ping", err)
	}

	if err != nil {
		log.Fatal("DB Connection Error", err)
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
			log.Fatal("InsertPunkIndex", err)
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
			log.Fatal("Latest Index filterCursor", err)
			return nil
		}
		var resultsFiltered []bson.M
		if err = filterCursor.All(ctx, &resultsFiltered); err != nil {
			log.Fatal("Latest Index resultsFiltered", err)
			return nil
		} else if len(resultsFiltered) == 0 {
			return nil
		}

		price := resultsFiltered[len(resultsFiltered)-1]["price"].(string)
		timestamp := resultsFiltered[len(resultsFiltered)-1]["timestamp"].(string)
		// unixTimestamp, _ := strconv.Atoi(resultsFiltered[len(resultsFiltered)-1]["timestamp"].(string))
		// timeT := time.Unix(int64(unixTimestamp), 0).UTC().String()

		values := map[string]interface{}{
			"price":     price,
			"timestamp": timestamp,
			// "timestampDate": timeT,
		}

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
			log.Fatal("Index history filterCursor", err)
			return nil
		}
		var resultsFiltered []bson.M
		if err = filterCursor.All(ctx, &resultsFiltered); err != nil {
			log.Fatal("Index history resultsFiltered", err)
			return nil
		}

		var values []map[string]interface{}
		dayCount := 0

		for _, result := range resultsFiltered {
			price := result["price"].(string)
			timestamp := result["timestamp"].(string)
			unixTimestamp, _ := strconv.Atoi(result["timestamp"].(string))
			timeT := time.Unix(int64(unixTimestamp), 0).UTC().String()

			if strings.Contains(timeT, "01:00") && dayCount < 30 {
				obj := map[string]interface{}{
					"price":     price,
					"timestamp": timestamp,
					// "timestampDate": timeT,
				}

				dayCount = dayCount + 1
				values = append([]map[string]interface{}{obj}, values...)
			}
		}

		return values
	} else {
		return nil
	}
}
