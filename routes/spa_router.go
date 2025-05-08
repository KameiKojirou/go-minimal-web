package routes

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)


func SPARouter(distDir, indexFile string) http.Handler {
	fs := http.FileServer(http.Dir(distDir))
  
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  // 1) Never serve /api
	  if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	  }
  
	  // 2) Compute the local file path
	  //    strip leading "/" so filepath.Join works correctly
	  relPath := strings.TrimPrefix(r.URL.Path, "/")
	  fullPath := filepath.Join(distDir, relPath)
  
	  // 3) If the file exists and is not a directory, serve it
	  if info, err := os.Stat(fullPath); err == nil && !info.IsDir() {
		// let the file server handle it
		fs.ServeHTTP(w, r)
		return
	  }
  
	  // 4) Fallback to index.html
	  http.ServeFile(w, r, filepath.Join(distDir, indexFile))
	})
  }