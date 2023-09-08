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
	//CHECK IF IS RUNNING ON DOCKER OR DEVELOPMENT NORMAL
	if os.Getenv("DOCKER_ENV") == "true" {
		MySQLStrConn = fmt.Sprintf("%s:%s@tcp(mysql_devbook:3306)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("MYSQL_ROOT"),
			os.Getenv("MYSQL_ROOT_PASSWORD"),
			os.Getenv("MYSQL_DATABASE"),
		)
	} else {
		MySQLStrConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("MYSQL_ROOT"),
			os.Getenv("MYSQL_ROOT_PASSWORD"),
			os.Getenv("MYSQL_DATABASE"),
		)

	}
}
