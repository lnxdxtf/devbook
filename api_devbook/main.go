package main

import (
	flag_mode "api/src"
	"api/src/config"
	"api/src/database"
	"api/src/server"
)

func init() {
	config.LoadEnv()
	db := database.MySQLDB{}
	db.Connect()
}

func main() {
	server := server.Server{}
	flag_mode.ProdMode()
	server.Start()
}
