package repository

import (
	"bytes"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/yuin/goldmark"
)

func InitDB(dsn string, seedOnEmpty bool) *sql.DB {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id BIGSERIAL PRIMARY KEY,
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

	if seedOnEmpty && count == 0 {
		seedFromMarkdown(db, filepath.Join("internal", "repository", "blog"))
	}

	return db
}

func seedFromMarkdown(db *sql.DB, dir string) {
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(d.Name()) != ".md" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("failed to read %s: %v", path, err)
			return nil
		}

		lines := strings.SplitN(string(data), "\n", 2)
		if len(lines) < 2 {
			log.Printf("skipping %s: not enough content", path)
			return nil
		}

		title := strings.TrimSpace(strings.TrimPrefix(lines[0], "#"))
		markdown := strings.TrimSpace(lines[1])

		var buf bytes.Buffer
		if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
			log.Printf("failed to convert markdown %s: %v", path, err)
			return nil
		}
		html := buf.String()

		if _, err := db.Exec(
			"INSERT INTO posts (title, content) VALUES ($1, $2)",
			title, html,
		); err != nil {
			log.Printf("failed to insert %s: %v", path, err)
		}
		return nil
	})
	if err != nil {
		log.Printf("error walking blog dir: %v", err)
	}
}
