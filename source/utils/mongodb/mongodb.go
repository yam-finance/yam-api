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
	//uri := "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"
	fmt.Println(uri)
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

func InsertAprYam(val map[string]interface{}) {
	if client != nil {
		result := AprYam{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database(dbName)
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

func InsertAprDegenerative(val map[string]interface{}) {
	if client != nil {
		result := AprDegenerative{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database(dbName)
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

// TODO: Add cycle
// TODO: Merge with InsertAssetIndex()
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

func InsertAssetIndex(val map[string]interface{}) {
	if client != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database(dbName)
		punkCollection := databaseRef.Collection("ustonks")

		/// @dev Insert new values
		_, err := punkCollection.InsertOne(ctx, bson.D{
			{Key: "cycle", Value: val["cycle"]},
			{Key: "price", Value: val["price"]},
			{Key: "timestamp", Value: val["timestamp"]},
		})
		if err != nil {
			log.Fatal("InsertAssetIndex", err)
		}
	}
}

func GetAprYam() map[string]interface{} {
	if client != nil {
		result := AprYam{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database(dbName)
		aprYamCollection := databaseRef.Collection("apryam")
		er := aprYamCollection.FindOne(ctx, bson.M{}).Decode(&result)
		if er != nil {
			fmt.Println(er)
			return nil
		}
		fmt.Println(result.Value)
		return result.Value
	}
	return nil
}

func GetAprDegenerative() map[string]interface{} {
	if client != nil {
		result := AprDegenerative{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database(dbName)
		aprYamDegenerativeCollection := databaseRef.Collection("aprdegenerative")
		er := aprYamDegenerativeCollection.FindOne(ctx, bson.M{}).Decode(&result)
		if er != nil {
			fmt.Println(er)
			return nil
		}
		fmt.Println(result.Value)
		return result.Value
	}
	return nil
}

// TODO: Merge with GetLatestAssetIndex()
func GetPunkIndex() map[string]interface{} {
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

// TODO: Merge with GetLatestAssetIndex()
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

		for i := len(resultsFiltered) - 1; i >= 0 && dayCount < 30; i-- {
			price := resultsFiltered[i]["price"].(string)
			timestamp := resultsFiltered[i]["timestamp"].(string)
			unixTimestamp, _ := strconv.Atoi(resultsFiltered[i]["timestamp"].(string))
			timeT := time.Unix(int64(unixTimestamp), 0).UTC().String()

			if strings.Contains(timeT, "01:0") {
				obj := map[string]interface{}{
					"price":     price,
					"timestamp": timestamp,
					// "timestampDate": timeT,
				}

				dayCount = dayCount + 1
				i--
				values = append([]map[string]interface{}{obj}, values...)
			}
		}

		return values
	} else {
		return nil
	}
}

func GetLatestAssetIndex(_cycle string) map[string]interface{} {
	if client != nil {
		// result := PunkIndex{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database(dbName)
		punkCollection := databaseRef.Collection("ustonks")

		// @dev Find a document that meets certain criteria
		// err := punkCollection.FindOne(ctx, bson.M{}).Decode(&result)
		// if err != nil {
		// 	log.Fatal(err)
		// 	return nil
		// }

		// @dev Find last document in collection
		filterCursor, err := punkCollection.Find(ctx, bson.M{"cycle": _cycle})
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

		cycle := resultsFiltered[len(resultsFiltered)-1]["cycle"].(string)
		price := resultsFiltered[len(resultsFiltered)-1]["price"].(string)
		timestamp := resultsFiltered[len(resultsFiltered)-1]["timestamp"].(string)
		unixTimestamp, _ := strconv.Atoi(resultsFiltered[len(resultsFiltered)-1]["timestamp"].(string))
		timeT := time.Unix(int64(unixTimestamp), 0).UTC().String()

		values := map[string]interface{}{
			"cycle":         cycle,
			"price":         price,
			"timestamp":     timestamp,
			"timestampDate": timeT,
		}

		return values
	} else {
		return nil
	}
}

func GetAssetIndexHistoryDaily(_cycle string) []map[string]interface{} {
	if client != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database(dbName)
		punkCollection := databaseRef.Collection("ustonks")

		// @dev Find last document in collection
		filterCursor, err := punkCollection.Find(ctx, bson.M{"cycle": _cycle})
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

		for i := len(resultsFiltered) - 1; i >= 0 && dayCount < 30; i-- {
			cycle := resultsFiltered[i]["cycle"].(string)
			price := resultsFiltered[i]["price"].(string)
			timestamp := resultsFiltered[i]["timestamp"].(string)
			unixTimestamp, _ := strconv.Atoi(resultsFiltered[i]["timestamp"].(string))
			timeT := time.Unix(int64(unixTimestamp), 0).UTC().String()

			if strings.Contains(timeT, "01:0") {
				obj := map[string]interface{}{
					"cycle":         cycle,
					"price":         price,
					"timestamp":     timestamp,
					"timestampDate": timeT,
				}

				dayCount = dayCount + 1
				i--
				values = append([]map[string]interface{}{obj}, values...)
			}
		}

		return values
	} else {
		return nil
	}
}
