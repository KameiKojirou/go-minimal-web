package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// Embed all files inside frontend/dist
//
//go:embed frontend/dist/*
var embeddedFiles embed.FS

// spaHandler serves files from the embedded filesystem,
// and falls back to index.html for client-side routes when
// a file is not found. It ignores /api requests.
type spaHandler struct {
	fileServer http.Handler
	indexPath  string
}

// NewSPAHandler creates a new spaHandler using the embedded files.
func NewSPAHandler() http.Handler {
	// Create a sub-filesystem starting at "frontend/dist" inside the embedded FS.
	distFS, err := fs.Sub(embeddedFiles, "frontend/dist")
	if err != nil {
		log.Fatalf("Failed to create sub filesystem: %v", err)
	}
	return spaHandler{
		fileServer: http.FileServer(http.FS(distFS)),
		indexPath:  "index.html",
	}
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Optionally ignore /api routes so that API requests can be handled elsewhere.
	if r.URL.Path == "/api" || strings.HasPrefix(r.URL.Path, "/api/") {
		http.NotFound(w, r)
		return
	}

	// Here we checkout if the file exists in the embedded FS.
	// The underlying http.FileServer does not provide an easy way to do this,
	// so we have to recreate this logic manually.
	path := r.URL.Path

	// Clean the path to remove any .. elements and a leading '/'
	path = filepath.Clean(path)
	if path == "/" {
		path = h.indexPath
	} else {
		// Remove the leading "/" to build our relative path.
		path = path[1:]
	}

	// Open the file from the embedded filesystem.
	f, err := embeddedFiles.Open(filepath.Join("frontend/dist", path))
	if err != nil {
		// If the file doesn't exist, fallback to index.html.
		http.ServeFile(w, r, filepath.Join("frontend/dist", h.indexPath))
		return
	}
	// Close the file once we're done.
	f.Close()

	// If it exists, use the embedded file server to serve the file.
	h.fileServer.ServeHTTP(w, r)
}

// SpaRoutes returns a ServeMux that serves the SPA from embedded files.
func SpaRoutes() *http.ServeMux {
	spaMux := http.NewServeMux()
	spaMux.Handle("/", NewSPAHandler())
	return spaMux
}