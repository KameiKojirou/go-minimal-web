package assets

import (
  "bytes"
  "embed"
  "io/fs"
  "net/http"
  "strings"
  "time"
)

//go:embed frontend/dist/*
//go:embed frontend/dist/**/*
var embeddedAssets embed.FS

// SPAHandler serves your embedded SPA, falling back to index.html.
func SPAHandler(indexFile string) http.Handler {
  // “frontend/dist” matches what you embedded above
  staticFS, err := fs.Sub(embeddedAssets, "frontend/dist")
  if err != nil {
    panic("embed fs.Sub failed: " + err.Error())
  }
  fileServer := http.FileServer(http.FS(staticFS))

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // 1) Don’t serve /api here
    if strings.HasPrefix(r.URL.Path, "/api") {
      http.NotFound(w, r)
      return
    }

    // 2) Clean up path
    reqPath := strings.TrimPrefix(r.URL.Path, "/")
    if reqPath == "" {
      reqPath = indexFile
    }

    // 3) If the file exists and is not a dir → serve it
    if info, err := fs.Stat(staticFS, reqPath); err == nil && !info.IsDir() {
      fileServer.ServeHTTP(w, r)
      return
    }

    // 4) Fallback → serve embedded index.html
    data, err := fs.ReadFile(staticFS, indexFile)
    if err != nil {
      http.NotFound(w, r)
      return
    }

    // bytes.Reader is an io.ReadSeeker
    reader := bytes.NewReader(data)
    // serve with a fresh modtime so browsers will revalidate
    http.ServeContent(w, r, indexFile, time.Now(), reader)
  })
}
