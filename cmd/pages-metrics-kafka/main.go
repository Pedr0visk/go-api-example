package main

import (
	"analytics/cmd/internal"
	internaldomain "analytics/internal"
	"analytics/internal/framework/envvar"
	"flag"
	"log"

	"go.uber.org/zap"
)

func main() {
	var env string

	flag.StringVar(&env, "env", "", "Environment Variables filename")
	flag.Parse()

	errC, err := run(env)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run(env string) (<-chan error, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "zap.NewProduction")
	}

	if err := envvar.Load(env); err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "envvar.Load")
	}

	vault, err := internal.NewVaultProvider()

	conf := envvar.New(vault)

	// -

	kafka, err := internal.NewKafkaConsumer(conf, "publisher-indexer")
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewKafkaConsumer")
	}
}
