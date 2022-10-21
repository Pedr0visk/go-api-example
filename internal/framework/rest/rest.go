package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hive-data-collector/internal/application/service"
	"hive-data-collector/internal/framework/kafka"
	"hive-data-collector/internal/framework/postgresql/db"
	"hive-data-collector/internal/framework/rest/middlewares"
	"log"
	"net/http"
)

func Run(configPath string) {
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

	fmt.Println("Go API REST Running on port " + "8000")
	fmt.Println("==================>")
	err := app.Run(":" + "8000")
	if err != nil {
		log.Fatalf("Error starting server: %v", err.Error())
	}
}
