package handlers

import (
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/NTsareva/url-shortener.git/cmd/storage"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Handlers struct {
	Chi     chi.Router
	Storage storage.Storage
}

func (h *Handlers) New() {
	h.Chi = chi.NewRouter()
	h.Storage.New()
}

func (h *Handlers) PostBodyHandler(res http.ResponseWriter, req *http.Request) {
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	req.Body.Close()

	res.Header().Set("Content-Type", "text/plain")

	if len(reqBody) == 0 {
		res.WriteHeader(400)
	}

	shortenedLink := randStringBytes(8)
	res.WriteHeader(201)
	h.Storage.Storage[shortenedLink] = string(reqBody)

	res.Write([]byte(shortenedLink))

}

func (h *Handlers) GetBodyHandler(res http.ResponseWriter, req *http.Request) {
	shortenedLink := chi.URLParam(req, "shortenLink")

	link, ok := h.Storage.Storage[shortenedLink]
	if !ok {
		res.WriteHeader(400)
	}
	res.WriteHeader(307)
	res.Write([]byte(link))
	res.Write([]byte("wowo"))

}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
