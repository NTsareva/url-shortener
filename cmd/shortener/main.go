package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/NTsareva/url-shortener.git/cmd/storage"
	"github.com/NTsareva/url-shortener.git/internal/server/handlers"
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
