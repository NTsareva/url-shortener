package main

import (
	"github.com/go-chi/chi/cmd/storage"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"math/rand"
	"net/http"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var linksStorage storage.Storage

func postBodyHandler(res http.ResponseWriter, req *http.Request) {
	linksStorage.New()

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	req.Body.Close()

	if len(reqBody) == 0 {
		res.WriteHeader(400)
	}

	shortenedLink := RandStringBytes(8)

	linksStorage.Storage[string(reqBody)] = shortenedLink

	res.Write([]byte(shortenedLink))

}

func getBodyHandler(res http.ResponseWriter, req *http.Request) {

}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func ShortenerRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", postBodyHandler)
	r.Get("/{shortenLink}", getBodyHandler)

	return r

}

func main() {

	http.ListenAndServe(":8080", ShortenerRouter())

}
