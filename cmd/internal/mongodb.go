package internal

import (
	"analytics/internal/framework/envvar"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewMongoDB(conf *envvar.Configuration) (*mongo.Client, error) {
	var database *mongo.Database

	get := func(v string) string {
		res, err := conf.Get(v)
		if err != nil {
			log.Fatalf("Couldn't get configuration value for %s: %s", v, err)
		}

		return res
	}

	// XXX: We will revisit this code in future episodes replacing it with another solution
	databaseHost := get("DATABASE_HOST")
	databasePort := get("DATABASE_PORT")
	//databaseUsername := get("DATABASE_USERNAME")
	//databasePassword := get("DATABASE_PASSWORD")
	databaseName := get("DATABASE_NAME")
	//databaseSSLMode := get("DATABASE_SSLMODE")
	// XXX: -

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", databaseHost, databasePort)))
	if err != nil {
		return nil, err
	}

	return client, err
	// Disconnect Mongo gracefully
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
}
