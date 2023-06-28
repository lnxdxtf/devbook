package main

import "api/src/server"

func main() {
	server := server.Server{}
	server.Start()
}
