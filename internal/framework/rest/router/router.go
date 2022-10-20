package router

import (
	"fmt"
	"hive-data-collector/internal/framework/rest/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
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

	// Routes
	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	app.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(http.StatusOK, "healthcheck")
	})

	return app
}
