package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"b-log.com/b-log/internal/handlers"
	"b-log.com/b-log/internal/repository"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	handlers.SetLayoutTemplate(tmpl)

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("missing required env var: DATABASE_URL")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing required env var: PORT")
	}
	addr := ":" + port

	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		log.Fatal("missing required env var: NATS_URL")
	}
	log.Println("Using NATS at", natsURL)

	seed := os.Getenv("SEED_ENABLED")
	if seed == "" {
		log.Fatal("missing required env var: SEED_ENABLED")
	}

	db := repository.InitDB(dbURL, seed == "true")
	defer db.Close()

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

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Println("Server starting on", addr)
	log.Fatal(srv.ListenAndServe())
}
