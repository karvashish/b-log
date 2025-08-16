package handlers

import (
	"net/http"
)

type RootHandler struct {
	standalone bool
}

func NewRootHandler(standalone bool) *RootHandler {
	return &RootHandler{standalone: standalone}
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var content string
	if h.standalone {
		content = `
	<div class="slider-wrapper">
		<button class="slider-btn" id="slide-left">&#9664;</button>
		<div id="posts" class="posts-container">
			<a class="post-link" href="/upload">
				<div class="post-box">
					<h2>Standalone mode</h2>
					<p>Posts are disabled. Use <strong>Upload</strong> to add content or run without STANDALONE.</p>
				</div>
			</a>
		</div>
		<button class="slider-btn" id="slide-right">&#9654;</button>
	</div>`
	} else {
		content = `
	<div class="slider-wrapper">
		<button class="slider-btn" id="slide-left">&#9664;</button>
		<div id="posts" class="posts-container" 
			hx-get="/posts" hx-trigger="load" hx-target="#posts" hx-swap="innerHTML">
		</div>
		<button class="slider-btn" id="slide-right">&#9654;</button>
	</div>`
	}

	renderLayout(w, "b-log", content)
}
