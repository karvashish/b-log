package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler { return &UploadHandler{} }

func (u *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.renderPage(w)
	case http.MethodPost:
		u.handleUpload(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (u *UploadHandler) renderPage(w http.ResponseWriter) {
	b, err := os.ReadFile(filepath.Join("templates", "upload.html"))
	if err != nil {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	renderLayout(w, "b-log - upload", string(b))
}

func (u *UploadHandler) handleUpload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 25<<20)
	if err := r.ParseMultipartForm(25 << 20); err != nil {
		http.Error(w, "invalid multipart form", http.StatusBadRequest)
		return
	}

	dir := filepath.Join("tmp", "uploads")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		http.Error(w, "failed to prepare storage", http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["files"]

	type result struct {
		Name string `json:"name"`
		Size int64  `json:"size"`
		Err  string `json:"err,omitempty"`
	}
	out := make([]result, 0, len(files))

	for _, fh := range files {
		fr, err := fh.Open()
		if err != nil {
			out = append(out, result{Name: fh.Filename, Err: "open failed"})
			continue
		}

		ts := time.Now().UnixNano()
		dstName := fmt.Sprintf("%d_%s", ts, filepath.Base(fh.Filename))
		dstPath := filepath.Join(dir, dstName)

		fd, err := os.Create(dstPath)
		if err != nil {
			_ = fr.Close()
			out = append(out, result{Name: fh.Filename, Err: "create failed"})
			continue
		}

		n, copyErr := io.Copy(fd, fr)
		_ = fr.Close()
		_ = fd.Close()
		if copyErr != nil {
			_ = os.Remove(dstPath)
			out = append(out, result{Name: fh.Filename, Err: "write failed"})
			continue
		}

		out = append(out, result{Name: fh.Filename, Size: n})
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"ok":    true,
		"files": out,
	})
}
