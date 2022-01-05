package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/create-go-app/chi-go-template/cmd"
	"github.com/create-go-app/chi-go-template/internal/config"
	"github.com/create-go-app/chi-go-template/internal/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Create router.
	r := chi.NewRouter()

	// Create config.
	c := config.NewConfig()

	// Set a logger middleware.
	r.Use(middleware.Logger)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(c.Server.ReadTimeout))

	// Get router with all routes.
	router.GetRoutes(r)

	// Run server instance.
	if err := cmd.Run(c, r); err != nil {
		log.Fatal(err)
		return
	}
}
