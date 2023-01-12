package internal

import (
	"analytics/internal"
	"analytics/internal/framework/envvar"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/url"
	"time"
)

func NewMongoDB(conf *envvar.Configuration) (*mongo.Client, error) {
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
	databaseSSLMode := get("DATABASE_SSLMODE")
	// XXX: -

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dsn := url.URL{
		Scheme: "mongodb",
		Host:   fmt.Sprintf("%s:%s", databaseHost, databasePort),
		Path:   databaseName,
	}

	q := dsn.Query()
	q.Add("ssl", databaseSSLMode)

	dsn.RawQuery = q.Encode()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn.String()))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "db.Ping")
	}

	return client, err
}
