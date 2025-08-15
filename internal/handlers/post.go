package handlers

import (
	"fmt"
	"html"
	"net/http"
	"strconv"

	"b-log.com/b-log/internal/repository"
)

type PostHandler struct {
	repo *repository.PostRepository
}

func NewPostHandler(repo *repository.PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := len(h.repo.GetAllPosts())
	offset := 0
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}
	if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
		offset = o
	}

	posts := h.repo.GetAllPosts()
	end := offset + limit
	if end > len(posts) {
		end = len(posts)
	}
	if offset >= len(posts) {
		return
	}

	renderPosts := func() string {
		htmlStr := ""
		for _, p := range posts[offset:end] {
			preview := p.Content
			if len(preview) > 500 {
				preview = preview[:500] + "..."
			}
			htmlStr += fmt.Sprintf(`
<a class="post-link" href="/post?id=%d">
	<div class="post-box">
		<h2>%s</h2>
		<p>%s</p>
	</div>
</a>
`, p.ID, html.EscapeString(p.Title), html.EscapeString(preview))
		}
		return htmlStr
	}

	if r.Header.Get("HX-Request") == "true" {
		fmt.Fprint(w, renderPosts())
	} else {

		content := `
<div class="slider-wrapper">
	<button class="slider-btn" id="slide-left">&#9664;</button>
	<div id="posts" class="posts-container">` + renderPosts() + `</div>
	<button class="slider-btn" id="slide-right">&#9654;</button>
</div>`
		renderLayout(w, "b-log", content)
	}
}

func (h *PostHandler) View(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	post := h.repo.GetPostByID(id)
	if post == nil {
		http.NotFound(w, r)
		return
	}

	renderPost := func() string {
		return fmt.Sprintf(`
<div class="post-box">
	<h2>%s</h2>
	<p>%s</p>
</div>
`, html.EscapeString(post.Title), html.EscapeString(post.Content))
	}

	if r.Header.Get("HX-Request") == "true" {
		fmt.Fprint(w, renderPost())
	} else {
		renderLayout(w, "b-log - "+html.EscapeString(post.Title), renderPost())
	}
}
