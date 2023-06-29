package server

import (
	"api/src/config"
	"api/src/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port    uint32
	address string
	router  *mux.Router
}

type ServerInterface interface {
	build()
	Start()
}

func (server *Server) build() {
	server.router = routers.Create()
	server.port = uint32(config.ApiPort)
	server.address = fmt.Sprintf("http://localhost:%d", server.port)
}
func (server *Server) Start() {
	server.build()
	msg := fmt.Sprintf("Server is running on %s", server.address)
	fmt.Println(msg)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", server.port), server.router))
}
