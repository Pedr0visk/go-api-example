package main

import (
	"analytics/cmd/internal"
	internaldomain "analytics/internal"
	"analytics/internal/application/service"
	"analytics/internal/framework/envvar"
	"analytics/internal/framework/kafka"
	"analytics/internal/framework/rest"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func main() {
	var env, address string

	flag.StringVar(&env, "env", "", "Environment Variables filename")
	flag.StringVar(&address, "address", ":9234", "HTTP Server Address")
	flag.Parse()

	errC, err := run(env, address)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run(env, address string) (<-chan error, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "zap.NewProduction")
	}

	if err := envvar.Load(env); err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "envvar.Load")
	}

	vault, err := internal.NewVaultProvider()
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewVaultProvider")
	}

	conf := envvar.New(vault)

	//-

	mongo, err := internal.NewMongoDB(conf)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewMongoDB")
	}

	kafka, err := internal.NewKafkaProducer(conf)
	if err != nil {
		return nil, fmt.Errorf("internal.NewKafka %w", err)
	}

	//-

	logging := func() gin.HandlerFunc {
		return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		})
	}

	//-

	srv, err := newServer(serverConfig{
		Address:     address,
		DB:          mongo,
		Kafka:       kafka,
		Middlewares: []func() gin.HandlerFunc{gin.Recovery, logging},
		Logger:      logger,
	})
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "newServer")
	}

	errC := make(chan error, 1)

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		logger.Info("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			_ = logger.Sync()

			err = mongo.Disconnect(ctxTimeout)
			stop()
			cancel()
			close(errC)
		}()

		if err := srv.Shutdown(ctxTimeout); err != nil { //nolint: contextcheck
			errC <- err
		}

		logger.Info("Shutdown completed")
	}()

	go func() {
		logger.Info("Listening and serving", zap.String("address", address))

		// "ListenAndServe always returns a non-nil error. After Shutdown or Close, the returned error is
		// ErrServerClosed."
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}

type serverConfig struct {
	Address     string
	DB          *mongo.Client
	Kafka       *internal.KafkaProducer
	Middlewares []func() gin.HandlerFunc
	Logger      *zap.Logger
}

func newServer(conf serverConfig) (*http.Server, error) {
	router := gin.New()

	for _, mw := range conf.Middlewares {
		router.Use(mw())
	}

	// -

	router.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(http.StatusOK, "healthcheck")
	})

	// -

	msgBroker := kafka.NewSpanMessageBroker(conf.Kafka.Producer, conf.Kafka.Topic)
	svc := service.NewSpanService(msgBroker)

	// -

	rest.NewSpanHandler(svc).Register(router)

	return &http.Server{
		Addr:    conf.Address,
		Handler: router,
	}, nil

}
