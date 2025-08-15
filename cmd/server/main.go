package main

import (
	"html/template"
	"log"
	"net/http"

	"b-log.com/b-log/internal/handlers"
	"b-log.com/b-log/internal/repository"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	handlers.SetLayoutTemplate(tmpl)

	db := repository.InitDB("tmp/b-log.db")
	postRepo := repository.NewPostRepository(db)

	rootHandler := handlers.NewRootHandler()
	postHandler := handlers.NewPostHandler(postRepo)
	healthHandler := handlers.NewHealthHandler()

	mux := http.NewServeMux()
	mux.Handle("/", rootHandler)
	mux.HandleFunc("/posts", postHandler.List)
	mux.HandleFunc("/post", postHandler.View)
	mux.HandleFunc("/healthz", healthHandler.Health)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
