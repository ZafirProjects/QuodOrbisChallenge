package server

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	NewServer := &Server{
		port: 3000,
	}

	middlewareStack := CreateStack(
		Logging,
	)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      middlewareStack(NewServer.RegisterRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
