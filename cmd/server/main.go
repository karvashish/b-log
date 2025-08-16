package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"b-log.com/b-log/internal/handlers"
	"b-log.com/b-log/internal/repository"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	handlers.SetLayoutTemplate(tmpl)

	standaloneEnv := os.Getenv("STANDALONE")
	standalone := true
	if standaloneEnv != "" {
		standalone = strings.ToLower(standaloneEnv) == "true"
	}
	if standalone {
		log.Println("Running in STANDALONE mode (no DB, no NATS)")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("missing required env var: PORT, defaulting to 8080")
		port = "8080"
	}
	addr := ":" + port

	var db *sql.DB
	if !standalone {
		dbURL := os.Getenv("DATABASE_URL")
		if dbURL == "" {
			log.Fatal("missing required env var: DATABASE_URL")
		}

		natsURL := os.Getenv("NATS_URL")
		if natsURL == "" {
			log.Fatal("missing required env var: NATS_URL")
		}
		log.Println("Using NATS at", natsURL)

		seed := os.Getenv("SEED_ENABLED")
		if seed == "" {
			log.Fatal("missing required env var: SEED_ENABLED")
		}

		db = repository.InitDB(dbURL, seed == "true")
		defer db.Close()
	}

	rootHandler := handlers.NewRootHandler(standalone)
	healthHandler := handlers.NewHealthHandler()
	uploadHandler := handlers.NewUploadHandler()

	mux := http.NewServeMux()
	mux.Handle("/", rootHandler)
	mux.HandleFunc("/healthz", healthHandler.Health)
	mux.HandleFunc("/upload", uploadHandler.ServeHTTP)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	if !standalone {
		postRepo := repository.NewPostRepository(db)
		postHandler := handlers.NewPostHandler(postRepo)
		mux.HandleFunc("/posts", postHandler.List)
		mux.HandleFunc("/post", postHandler.View)
	} else {
		log.Println("STANDALONE: skipping /posts and /post routes")
	}

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Println("Server starting on", addr)
	log.Fatal(srv.ListenAndServe())
}
