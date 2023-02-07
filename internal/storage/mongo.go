package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	productCol      *mongo.Collection
	crawlThreadCol  *mongo.Collection
	crawlKeywordCol *mongo.Collection
	crawlProductCol *mongo.Collection
	userCol         *mongo.Collection
	userLikeCol     *mongo.Collection
	orderCol        *mongo.Collection
	notificationCol *mongo.Collection
}

func NewMongoDB() *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	mongoClient, err := makeMongoClient(ctx)
	checkErr(err, "Connection in mongodb")
	// checkErr(mongoClient.Ping(ctx, readpref.Primary()), "Ping error in mongoconnect")
	db := mongoClient.Database(viper.GetString("MONGO_DB_NAME"))

	return &MongoDB{
		productCol:      db.Collection("products"),
		crawlThreadCol:  db.Collection("crawl_threads"),
		crawlKeywordCol: db.Collection("crawl_keywords"),
		crawlProductCol: db.Collection("crawl_products"),
		userCol:         db.Collection("users"),
		userLikeCol:     db.Collection("user_likes"),
		orderCol:        db.Collection("orders"),
		notificationCol: db.Collection("notifications"),
	}
}

func makeMongoClient(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + viper.GetString("MONGO_URL") + "/" + viper.GetString("MONGO_DB_NAME") + "?&connect=direct&replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false").SetAuth(options.Credential{
		Username: viper.GetString("MONGO_USERNAME"),
		Password: viper.GetString("MONGO_PASSWORD"),
	})
	mongoClient, err := mongo.Connect(ctx, clientOptions)

	return mongoClient, err
}

func checkErr(err error, location string) {
	if err != nil {
		fmt.Println("Error occured: " + location)
		fmt.Println("Message: " + err.Error())
	}
}
