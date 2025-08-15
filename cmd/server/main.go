package main

import (
	"log"
	"net/http"

	"b-log.com/b-log/internal/handlers"
	"b-log.com/b-log/internal/repository"
)

func main() {

	postRepo := repository.NewPostRepository()

	rootHandler := handlers.NewRootHandler()
	postHandler := handlers.NewPostHandler(postRepo)
	healthHandler := handlers.NewHealthHandler()

	mux := http.NewServeMux()
	mux.Handle("/", rootHandler)
	mux.HandleFunc("/posts", postHandler.List)
	mux.HandleFunc("/post", postHandler.View)
	mux.HandleFunc("/healthz", healthHandler.Health)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
