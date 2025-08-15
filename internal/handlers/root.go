package handlers

import "net/http"

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content := `
	<div class="slider-wrapper">
			<button class="slider-btn" id="slide-left">&#9664;</button>
			<div id="posts" class="posts-container" 
				hx-get="/posts" hx-trigger="load" hx-target="#posts" hx-swap="innerHTML">
			</div>
			<button class="slider-btn" id="slide-right">&#9654;</button>
	</div>
`
	renderLayout(w, "b-log", content)
}
