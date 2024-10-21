package main

import (
	"log"

	"github.com/banhawy/go-microservices/internal/database"
	"github.com/banhawy/go-microservices/internal/server"
)

func main() {
	db, err := database.NewdatabaseClient()
	if err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}
