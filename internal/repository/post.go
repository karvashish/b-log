package repository

import (
	"database/sql"
)

type Post struct {
	ID      int
	Title   string
	Content string
}

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) GetAllPosts() []Post {
	rows, err := r.db.Query("SELECT id, title, content FROM posts ORDER BY id ASC")
	if err != nil {
		return []Post{}
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content); err == nil {
			posts = append(posts, p)
		}
	}
	return posts
}

func (r *PostRepository) GetPostByID(id int) *Post {
	var p Post
	err := r.db.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", id).Scan(&p.ID, &p.Title, &p.Content)
	if err != nil {
		return nil
	}
	return &p
}
