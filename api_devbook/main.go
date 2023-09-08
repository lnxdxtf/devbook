package main

import (
	flag_mode "api/src"
	"api/src/config"
	"api/src/database"
	"api/src/server"
	"time"
)

func init() {
	config.LoadEnv()
	db := database.MySQLDB{}
	// wait for mysql container to be ready
	time.Sleep(30 * time.Second)
	db.Connect()
}

func main() {
	server := server.Server{}
	flag_mode.ProdMode()
	server.Start()
}
