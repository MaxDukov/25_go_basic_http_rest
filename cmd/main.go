package main

import (
	"log"
	"os"

	"local/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "HTTP-SERVER: ", log.LstdFlags|log.Lmsgprefix)

	srv := server.New(logger)

	logger.Println("Starting server on :8080")

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
