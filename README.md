b-log
=====

A blog application written in Go.  

Features
--------
- Go-powered backend using net/http
- HTMX integration for dynamic content loading

Project Structure
-----------------
/cmd           - Application entry points
/internal
  /handlers    - HTTP handlers for listing and viewing posts
  /repository  - In-memory post repository
/static        - Static assets (CSS, JS, etc.)

Getting Started
---------------
Prerequisites:
- Go 1.20+

Installation:
1. Clone the repository:
   git clone git@github.com:<your-username>/b-log.git
2. Change into the project directory:
   cd b-log
3. Run the server:
   go run ./cmd/server

Usage:
- Visit http://localhost:8080 to view the blog
- Click a post to view its full content
- Use 'limit' and 'offset' query parameters for pagination:
  http://localhost:8080/posts?limit=2&offset=2

Development:
- Posts are stored in-memory via PostRepository
- Edit internal/repository/posts.go to modify seeded content


License
-------
MIT
