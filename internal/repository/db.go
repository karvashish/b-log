package repository

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL
	);`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count); err != nil {
		log.Fatalf("failed to count posts: %v", err)
	}

	if count == 0 {
		seedPosts := []Post{
			{Title: "First Beginnings", Content: "This is the very first blog post in our series, marking the start of our journey..."},
		}
		for _, p := range seedPosts {
			_, err := db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", p.Title, p.Content)
			if err != nil {
				log.Printf("failed to insert seed post: %v", err)
			}
		}
	}

	return db
}
