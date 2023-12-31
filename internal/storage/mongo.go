package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	productCol        *mongo.Collection
	crawlThreadCol    *mongo.Collection
	crawlKeywordCol   *mongo.Collection
	crawlProductCol   *mongo.Collection
	userCol           *mongo.Collection
	userLikeCol       *mongo.Collection
	orderCol          *mongo.Collection
	notificationCol   *mongo.Collection
	keywordCol        *mongo.Collection
	keywordProductCol *mongo.Collection
}

func NewMongoDB() *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	mongoClient, err := makeMongoClient(ctx)
	checkErr(err, "Connection in mongodb")
	// checkErr(mongoClient.Ping(ctx, readpref.Primary()), "Ping error in mongoconnect")
	db := mongoClient.Database(viper.GetString("MONGO_DB_NAME"))

	return &MongoDB{
		productCol:        db.Collection("products"),
		crawlThreadCol:    db.Collection("crawl_threads"),
		crawlKeywordCol:   db.Collection("crawl_keywords"),
		crawlProductCol:   db.Collection("crawl_products"),
		userCol:           db.Collection("users"),
		userLikeCol:       db.Collection("user_likes"),
		notificationCol:   db.Collection("notifications"),
		orderCol:          db.Collection("orders"),
		keywordCol:        db.Collection("keywords"),
		keywordProductCol: db.Collection("keyword_products"),
	}
}

func makeMongoClient(ctx context.Context) (*mongo.Client, error) {
	localUri := "mongodb://" + viper.GetString("MONGO_USERNAME") + ":" + viper.GetString("MONGO_PASSWORD") + "@" + viper.GetString("MONGO_URL")
	clientOptions := options.Client().ApplyURI(localUri)

	// Connect to MongoDB
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("hoit1")
		fmt.Println("Error occured: " + err.Error())
	}
	return mongoClient, err
}

func checkErr(err error, location string) {
	if err != nil {
		log.Println("hoit2")
		fmt.Println("Error occured: " + location)
		fmt.Println("Message: " + err.Error())
	}
}
