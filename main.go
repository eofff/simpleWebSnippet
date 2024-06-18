package main

import (
	"database/sql"
	"fmt"
	"gosample/migrations"
	"gosample/redisApi"
	"gosample/repository"
	"gosample/services"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	fmt.Println(psqlconn)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	migrations.Migrate(db)
	repository.Init(db)

	result, err := services.SampleGetById(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	resultB, err := redisApi.Exists("kek")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resultB)

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
