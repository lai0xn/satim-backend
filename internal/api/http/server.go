package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/laix0n/satim/config"
	"github.com/laix0n/satim/internal/api/http/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() (*mongo.Client, context.Context, error) {
	config.LoadENV()
	mongoUrl := config.Dburl
	fmt.Println("Connecting to MongoDB...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(mongoUrl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		cancel()
		return nil, nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		cancel()
		return nil, nil, fmt.Errorf("could not connect to MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB!")
	return client, ctx, nil
}

func main() {
	client, ctx, err := DB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})


	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to Test API 👋"))
		})
		r.Post("/test", handlers.SendUrl) 
	})

	log.Println("Server starting on :5000...")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
