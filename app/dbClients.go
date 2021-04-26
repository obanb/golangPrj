package app

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func getMySqlClient(uri string) *sqlx.DB {
	client, err := sqlx.Open("mysql", uri)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}

func getMongoClient(uri string) *mongo.Client {
	defaultCancellationCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	client, err := mongo.Connect(defaultCancellationCtx, options.Client().ApplyURI(uri))

	//defer func() {
	//	if err = client.Disconnect(defaultCancellationCtx); err != nil {
	//		panic(err)
	//		fmt.Print("disconnect")
	//
	//		cancel()
	//	}
	//}()
	if err != nil {
		cancel()
	}

	fmt.Print("conn")
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		fmt.Print("conn err")

		log.Fatal(err)
	}
	return client
}
