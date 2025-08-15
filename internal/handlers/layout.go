package handlers

import (
	"fmt"
	"net/http"
)

func renderLayout(w http.ResponseWriter, title, content string) {
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>%s</title>
	<script src="https://unpkg.com/htmx.org@1.9.10"></script>
	<style>
		html, body { height: 100%%; margin: 0; }
		body {
			background-color: #121212;
			color: #e0e0e0;
			font-family: 'Segoe UI', Roboto, sans-serif;
			line-height: 1.6;
			display: flex;
			flex-direction: column;
			min-height: 100vh;
		}
		header {
			background-color: #1e1e1e;
			padding: 1rem;
			font-size: 1.8rem;
			font-weight: bold;
			border-bottom: 1px solid #333;
			text-align: left;
			letter-spacing: 1px;
		}
		main {
			max-width: 100%%;
			padding: 2rem;
			box-sizing: border-box;
		}

		.slider-wrapper {
			position: relative;
			overflow: hidden;
		}
		.posts-container {
			display: flex;
			gap: 1.5rem;
			transition: transform 0.4s ease;
			will-change: transform;
		}
		.post-link { text-decoration: none; color: inherit; flex: 0 0 50rem; margin-block: 12rem }

		.post-box {
			background: linear-gradient(180deg, #1f1f1f 0%%, #171717 100%%);
			padding: 1.25rem;
			border: 1px solid #2a2a2a;
			border-radius: 8px;
			display: flex;
			flex-direction: column;
			justify-content: space-between;
			box-shadow: 0 8px 16px rgba(0, 0, 0, 0.6), inset 0 1px 0 rgba(255, 255, 255, 0.02);
			transition: transform 0.2s ease, box-shadow 0.2s ease, background-color 0.2s ease, border-color 0.2s ease;
			cursor: pointer;
		}
		.post-box:hover {
			transform: translateY(-4px);
			box-shadow: 0 12px 24px rgba(0, 0, 0, 0.65);
			background-color: #232323;
			border-color: #3a3a3a;
		}
		.post-box h2 {
			margin: 0 0 0.6rem 0;
			font-size: 1.4rem;
			font-weight: 700;
			letter-spacing: 0.2px;
			color: #ffffff;
		}
		.post-box p {
			flex-grow: 1;
			color: #c6c6c6;
			line-height: 1.55;
			margin: 0;
		}

		.slider-btn {
			position: absolute;
			top: 50%%;
			transform: translateY(-50%%);
			background-color: rgba(0,0,0,0.5);
			color: white;
			border: none;
			font-size: 2rem;
			padding: 0.4rem 0.8rem;
			cursor: pointer;
			z-index: 2;
		}
		.slider-btn:hover {
			background-color: rgba(0,0,0,0.8);
		}
		#slide-left { left: 0; }
		#slide-right { right: 0; }
	</style>
</head>
<body>
	<header>
		<a href="/" style="color:inherit; text-decoration:none;">b-log</a>
	</header>
	<main>
		<div class="slider-wrapper">
			<button class="slider-btn" id="slide-left">&#9664;</button>
			<div class="posts-container">%s</div>
			<button class="slider-btn" id="slide-right">&#9654;</button>
		</div>
	</main>
	<script>
		(function(){
			const container = document.querySelector('.posts-container');
			const step = 320; // pixels per slide
			let offset = 0;

			document.getElementById('slide-left').addEventListener('click', () => {
				offset = Math.min(offset + step, 0);
				container.style.transform = 'translateX(' + offset + 'px)';
			});
			document.getElementById('slide-right').addEventListener('click', () => {
				const maxOffset = -(container.scrollWidth - container.clientWidth);
				offset = Math.max(offset - step, maxOffset);
				container.style.transform = 'translateX(' + offset + 'px)';
			});
		})();
	</script>
</body>
</html>`, title, content)
}
