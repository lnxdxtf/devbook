package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MySQLStrConn = ""
	ApiPort      = 0000
)

func LoadEnv() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ApiPort = 8000
	}

	MySQLStrConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_MYSQL_API_USER"),
		os.Getenv("DB_MYSQL_API_PASSWORD"),
		os.Getenv("DB_MYSQL_API_DATABASE_NAME"),
	)
}
