package rest

import (
	"analytics/internal"
	"analytics/internal/application/service"
	"analytics/internal/framework/kafka"
	"analytics/internal/framework/mongodb/db"
	"analytics/internal/framework/rest/middlewares"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"log"
	"net/http"
)

const otelName = "github.com/MarioCarrion/todo-rest/internal/rest"

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Error       string            `json:"error"`
	Validations validation.Errors `json:"validations,omitempty"`
}

func renderErrorResponse(c *gin.Context, msg string, err error) {
	resp := ErrorResponse{Error: msg}
	status := http.StatusInternalServerError

	var ierr *internal.Error
	if !errors.As(err, &ierr) {
		resp.Error = "internal error"
	} else {
		switch ierr.Code() {
		case internal.ErrorCodeNotFound:
			status = http.StatusNotFound
		case internal.ErrorCodeInvalidArgument:
			status = http.StatusBadRequest

			var verrors validation.Errors
			if errors.As(ierr, &verrors) {
				resp.Validations = verrors
			}
		case internal.ErrorCodeUnknown:
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

	prepo := db.NewPageRepository()
	pmsgRepo := kafka.NewTraceMessageBroker()
	psvc := service.NewPageService(prepo, pmsgRepo)

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

	// ================== Span Routes
	traceHandler := NewPageHandler(psvc)
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
