package main

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	HTTPPort      string
	DbHost        string
	DbPort        int
	DbUser        string
	DbPassword    string
	DbName        string
	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisDb       int
}

func (c *Config) Load() {
	var exists bool
	var err error
	c.DbHost, exists = os.LookupEnv("DBHOST")
	if !exists {
		log.Fatal("there is no DBHOST var")
	}

	port, exists := os.LookupEnv("DBPORT")
	if !exists {
		log.Fatal("there is no DBPORT var")
	}

	c.DbPort, err = strconv.Atoi(port)
	if err != nil {
		log.Fatal("DBPORT var is incorrect")
	}

	c.DbUser, exists = os.LookupEnv("DBUSER")
	if !exists {
		log.Fatal("there is no DBUSER var")
	}

	c.DbPassword, exists = os.LookupEnv("DBPASSWORD")
	if !exists {
		log.Fatal("there is no DBPASSWORD var")
	}

	c.DbName, exists = os.LookupEnv("DBNAME")
	if !exists {
		log.Fatal("there is no DBNAME var")
	}

	c.RedisHost, exists = os.LookupEnv("REDISHOST")
	if !exists {
		log.Fatal("there is no REDISHOST var")
	}

	redisPort, exists := os.LookupEnv("REDISPORT")
	if !exists {
		log.Fatal("there is no REDISPORT var")
	}

	c.RedisPort, err = strconv.Atoi(redisPort)
	if err != nil {
		log.Fatal("REDISPORT var incorrect")
	}

	c.RedisPassword, exists = os.LookupEnv("REDISPASSWORD")
	if !exists {
		log.Fatal("there is no REDISPASSWORD var")
	}

	redisDb, exists := os.LookupEnv("REDISDB")
	if !exists {
		log.Fatal("there is no REDISDB var")
	}

	c.RedisDb, err = strconv.Atoi(redisDb)
	if err != nil {
		log.Fatal("REDISDB var incorrect")
	}

	c.HTTPPort, exists = os.LookupEnv("HTTP_PORT")
	if !exists {
		log.Fatal("there is no HTTP_PORT var")
	}
}
