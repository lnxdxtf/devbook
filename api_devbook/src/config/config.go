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
	ApiPort      = 8000
	JwtSecretKey []byte // For signing JWT tokens
)

func LoadEnv() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))

	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	if err != nil {
		ApiPort = 8000
	}

	MySQLStrConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_MYSQL_API_USER"),
		os.Getenv("DB_MYSQL_API_PASSWORD"),
		os.Getenv("DB_MYSQL_API_DATABASE_NAME"),
	)
}
