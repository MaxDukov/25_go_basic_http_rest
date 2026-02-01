package server

import (
	"log"
	"net/http"
	"time"

	"local/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func New(logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.RootHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: httpServer,
	}
}
