package rest

import (
	"fmt"
	"hive-data-collector/internal/framework/rest/router"
	"log"
)

func Run(configPath string) {
	if configPath == "" {
		configPath = "modules/advertisements/data/config.yml"
	}

	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	err := web.Run(":" + conf.Server.Port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err.Error())
	}
}
