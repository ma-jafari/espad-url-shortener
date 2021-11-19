package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const (
	DB_CONNECTION_STRING = "mongodb://localhost:27017"

	DATABASE_NAME         = "url-shortener"
	URLS_COLLECTION_NAME  = "urls"
	USERS_COLLECTION_NAME = "users"
)

var db *mongo.Database

func OpenConnection() {
	i := 0
	for {
		if client, err := mongo.NewClient(options.Client().ApplyURI(DB_CONNECTION_STRING)); err != nil {
			fmt.Println(err.Error())
		} else {
			ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
			err = client.Connect(ctx)
			if err != nil {
				log.Fatal(err)
			}

			err = client.Ping(ctx, readpref.Primary())
			if err != nil {
				fmt.Println(err.Error())
			}
			db = client.Database(DATABASE_NAME)
			break
		}

		i++
		if i > 600 {
			break
		}
		time.Sleep(time.Second)
	}
}

func GetUrlsCollection() (serviceCollection *mongo.Collection, dbContext context.Context, cancelFunc context.CancelFunc) {
	serviceCollection = db.Collection(URLS_COLLECTION_NAME)
	dbContext, cancelFunc = context.WithTimeout(context.Background(), 20*time.Minute)
	return
}

func GetUsersCollection() (serviceCollection *mongo.Collection, dbContext context.Context, cancelFunc context.CancelFunc) {
	serviceCollection = db.Collection(USERS_COLLECTION_NAME)
	dbContext, cancelFunc = context.WithTimeout(context.Background(), 20*time.Minute)
	return
}
