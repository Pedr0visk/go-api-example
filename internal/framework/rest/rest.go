package rest

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"hive-data-collector/internal"
	"hive-data-collector/internal/application/service"
	"hive-data-collector/internal/domain"
	"hive-data-collector/internal/framework/kafka"
	"hive-data-collector/internal/framework/postgresql/db"
	"hive-data-collector/internal/framework/rest/middlewares"
	"log"
	"net/http"
)

const otelName = "github.com/MarioCarrion/todo-api/internal/rest"

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Error       string            `json:"error"`
	Validations validation.Errors `json:"validations,omitempty"`
}

func renderErrorResponse(c *gin.Context, msg string, err error) {
	resp := ErrorResponse{Error: msg}
	status := http.StatusInternalServerError

	var ierr *domain.Error
	if !errors.As(err, &ierr) {
		resp.Error = "internal error"
	} else {
		switch ierr.Code() {
		case domain.ErrorCodeNotFound:
			status = http.StatusNotFound
		case domain.ErrorCodeInvalidArgument:
			status = http.StatusBadRequest

			var verrors validation.Errors
			if errors.As(ierr, &verrors) {
				resp.Validations = verrors
			}
		case domain.ErrorCodeUnknown:
			fallthrough
		default:
			status = http.StatusInternalServerError
		}
	}

	if err != nil {
		//_, span := otel.Tracer(otelName).Start(c, "renderErrorResponse")
		//defer span.End()
		//
		//span.RecordError(err)
		fmt.Println("otel called...")
	}

	fmt.Printf("Error: %v\n", err)

	c.JSON(status, resp)
}

func setConfig(configPath string) {
	internal.Setup(configPath)
	db.SetupDB()
	gin.SetMode(internal.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}

	setConfig(configPath)
	conf := internal.GetConfig()

	// -

	trepo := db.NewTraceRepository()
	tmsgRepo := kafka.NewTraceMessageBroker()
	tsvc := service.NewTraceService(trepo, tmsgRepo)

	// -

	app := gin.New()

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
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
	}))

	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	// app.NoRoute(middlewares.NoRouteHandler()) TODO: create no router handler

	// ================== Trace Routes
	traceHandler := NewTraceHandler(tsvc)
	traceHandler.Register(app)

	// ================== Docs Routes
	// app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//
	app.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(http.StatusOK, "healthcheck")
	})

	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	err := app.Run(":" + conf.Server.Port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err.Error())
	}
}
