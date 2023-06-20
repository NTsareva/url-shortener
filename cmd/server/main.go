package main

import (
	"net/http"

	"github.com/go-chi/chi/cmd/storage"
	"github.com/go-chi/chi/internal/server/handlers"
	"github.com/go-chi/chi/v5"
)

var linksStorage storage.Storage

func ShortenerRouter() chi.Router {
	var handlers handlers.Handlers
	handlers.New()

	r := handlers.Chi

	r.Post("/", handlers.PostBodyHandler)
	r.Get("/{shortenLink}", handlers.GetBodyHandler)

	return r
}

func main() {
	http.ListenAndServe(":8080", ShortenerRouter())
}
