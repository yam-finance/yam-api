package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

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
	uri := "mongodb://localhost:27017/ugasmedian?readPreference=primary&appname=MongoDB%20Compass&ssl=false"
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
	//defer client.Disconnect(ctx)
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
		punkCollection := databaseRef.Collection("apryam")
		// result := punkCollection.FindOne(context.TODO(), bson.D{})
		//punkCollection.FindOneAndUpdate()
		update := bson.M{
			"$set": bson.M{"apryam": val},
		}
		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}
		er := punkCollection.FindOneAndUpdate(ctx, bson.M{}, update, &opt).Decode(&result)
		if er != nil {
			fmt.Println(er)
			return
		}

		//	fmt.Println(result.Value)
		// uPunksResult, err := punkCollection.InsertOne(ctx, bson.D{
		// 	{Key: "apryam", Value: val},
		// })

		// fmt.Println(uPunksResult)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}

}
func InsertAprDegenerative(val map[string]float64) {
	if client != nil {
		result := AprDegenerative{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database("ugasmedian")
		punkCollection := databaseRef.Collection("aprdegenerative")
		// result := punkCollection.FindOne(context.TODO(), bson.D{})
		//punkCollection.FindOneAndUpdate()
		update := bson.M{
			"$set": bson.M{"aprdegenerative": val},
		}
		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}
		er := punkCollection.FindOneAndUpdate(ctx, bson.M{}, update, &opt).Decode(&result)
		if er != nil {
			fmt.Println(er)
			return
		}

		//	fmt.Println(result.Value)
		// uPunksResult, err := punkCollection.InsertOne(ctx, bson.D{
		// 	{Key: "apryam", Value: val},
		// })

		// fmt.Println(uPunksResult)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}

}
func GetAprYam() float64 {
	if client != nil {
		result := AprYam{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database("ugasmedian")
		punkCollection := databaseRef.Collection("apryam")
		// result := punkCollection.FindOne(context.TODO(), bson.D{})
		//punkCollection.FindOneAndUpdate()

		er := punkCollection.FindOne(ctx, bson.M{}).Decode(&result)
		if er != nil {
			fmt.Println(er)
			return 0
		}

		fmt.Println(result.Value)
		// uPunksResult, err := punkCollection.InsertOne(ctx, bson.D{
		// 	{Key: "apryam", Value: val},
		// })

		// fmt.Println(uPunksResult)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		return result.Value
	}
	return 0

}
func GetAprDegenerative() map[string]float64 {
	if client != nil {
		result := AprDegenerative{}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		databaseRef := client.Database("ugasmedian")
		punkCollection := databaseRef.Collection("aprdegenerative")
		// result := punkCollection.FindOne(context.TODO(), bson.D{})
		//punkCollection.FindOneAndUpdate()

		er := punkCollection.FindOne(ctx, bson.M{}).Decode(&result)
		if er != nil {
			fmt.Println(er)
			return map[string]float64{"MAR21": 0, "JUN21": 0}
		}

		fmt.Println(result.Value)
		// uPunksResult, err := punkCollection.InsertOne(ctx, bson.D{
		// 	{Key: "apryam", Value: val},
		// })

		// fmt.Println(uPunksResult)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		return result.Value
	}
	return map[string]float64{"MAR21": 0, "JUN21": 0}

}
