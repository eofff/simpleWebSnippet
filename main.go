package main

import (
	"database/sql"
	"fmt"
	"gosample/migrations"
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
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
