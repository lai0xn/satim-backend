package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)



func Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	log.Println("Starting server on :8080")

	http.ListenAndServe(":8080", r)
}
