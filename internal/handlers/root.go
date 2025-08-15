package handlers

import "net/http"

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content := `
<div id="posts" class="posts-container" 
	hx-get="/posts" hx-trigger="load" hx-target="#posts" hx-swap="innerHTML">
</div>
`
	renderLayout(w, "b-log", content)
}
