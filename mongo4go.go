package mongo4go

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func InitMongoServer(username, password, server string, port int) (*mongo.Client, context.Context) {
	credential := options.Credential{
		Username: username,
		Password: password,
	}

	uri := "mongodb://" + server + ":" + fmt.Sprint(port)
	log.Printf("Trying to connect to <%s>\n", uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	log.Printf("Pinging <%s>\n", uri)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	log.Printf("Successfully connected to: <%s>\n", uri)
	return client, ctx
}
