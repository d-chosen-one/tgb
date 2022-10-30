package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var ClientMongo *mongo.Client

func InitMongo(url string) error {
	err := connect(url)
	return err
}

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// This is a user defined method that returns db.Client,
// context.Context, context.CancelFunc and error.
// db.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.

func connect(uri string) error {
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		return err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err == nil {
		ClientMongo = client
	}

	defer cancelFunc()
	return ping(ClientMongo, ctx)
}

// This is a user defined method that accepts
// db.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error {

	// db.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	log.Println("connected to Mongo successfully")
	return nil
}
