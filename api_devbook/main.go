package main

import (
	"api/src/config"
	"api/src/database"
	"api/src/server"
)

func main() {
	config.LoadEnv()
	db := database.MySQLDB{}
	db.Connect()
	server := server.Server{}
	server.Start()
}
