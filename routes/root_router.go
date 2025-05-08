package routes

import (
  "net/http"
)

func RootRouter() {
  rootMux := http.NewServeMux()

  rootMux.Handle("/api/", http.StripPrefix("/api", APIRouter()))
  rootMux.Handle("/", SPARouter("frontend/dist", "index.html"))

  // 6) Start your server with the single handler:
  http.ListenAndServe(":1323", rootMux)
}
