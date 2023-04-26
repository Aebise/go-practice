package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Ctx context.Context
var err error

func ConnectDB() (*mongo.Client, context.Context) {
	fmt.Println("Connecting to DB...")
	Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	Client, err = mongo.Connect(Ctx, options.Client().ApplyURI("mongodb+srv://ebisechimdessa:P9L36w1gtleqev8h@cluster0.hq8er8b.mongodb.net/?retryWrites=true&w=majority"))
	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://ebisechimdessa:P9L36w1gtleqev8h@cluster0.hq8er8b.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println("error connecting to db: ", err)
		// panic(err)
	}

	// defer Client.Disconnect(Ctx)

	return Client, Ctx
}
