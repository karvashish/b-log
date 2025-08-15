package repository

import (
	"database/sql"
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"path/filepath"

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
		seedFromJSON(db, filepath.Join("internal", "repository", "blog"))
	}

	return db
}

func seedFromJSON(db *sql.DB, dir string) {
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(d.Name()) != ".json" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("failed to read %s: %v", path, err)
			return nil
		}

		var post Post
		if err := json.Unmarshal(data, &post); err != nil {
			log.Printf("failed to parse %s: %v", path, err)
			return nil
		}

		if post.Title == "" || post.Content == "" {
			log.Printf("skipping %s: missing title or content", path)
			return nil
		}

		_, err = db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", post.Title, post.Content)
		if err != nil {
			log.Printf("failed to insert %s: %v", path, err)
		}
		return nil
	})
	if err != nil {
		log.Printf("error walking blog dir: %v", err)
	}
}
