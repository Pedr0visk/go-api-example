package db

import (
	"fmt"
	"hive-data-collector/internal/framework/postgresql/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() {
	var db = DB

	configuration := config.GetConfig()

	driver := configuration.Database.Driver
	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	if driver == "sqlite" { // SQLITE
		db, err = gorm.Open(sqlite.Open("modules/advertisements/"+database+".db"), &gorm.Config{})
		if err != nil {
			fmt.Println("db err: ", err)
		}
	} else if driver == "postgres" { // POSTGRES
		db, err = gorm.Open(postgres.Open("host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password), &gorm.Config{})
		if err != nil {
			fmt.Println("db err: ", err)
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting the sqlDB %v\n", err.Error())
	}

	// Change this to true if you want to see SQL queries
	sqlDB.SetMaxIdleConns(configuration.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(configuration.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)

	DB = db

	migration()
}

// Auto migrate project models
func migration() {
	DB.AutoMigrate(&Track{})
}

func GetDB() *gorm.DB {
	return DB
}
