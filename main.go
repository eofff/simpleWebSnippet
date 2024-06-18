package main

import (
	"database/sql"
	"fmt"
	"gosample/controllers"
	"gosample/migrations"
	"gosample/redisApi"
	"gosample/repository"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var config Config

func main() {
	_, exists := os.LookupEnv("PROD")

	if !exists {
		err := godotenv.Load("./dev.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	config.Load()

	err := redisApi.Init(config.RedisHost, config.RedisPort, config.RedisPassword, config.RedisDb)
	if err != nil {
		log.Fatal(err)
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	err = db.Ping()
	CheckError(err)

	migrations.Migrate(db)
	repository.Init(db)

	e := echo.New()
	e.GET("/", controllers.IndexGetAll)
	e.GET("/:id", controllers.IndexGetById)
	e.POST("/", controllers.IndexPost)
	e.Logger.Fatal(e.Start(":" + config.HTTPPort))
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
