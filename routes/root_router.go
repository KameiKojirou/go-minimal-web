package routes

import (
  "net/http"
  "main/assets"
)

func RootRouter() {
  rootMux := http.NewServeMux()

  rootMux.Handle("/api/", http.StripPrefix("/api", APIRouter()))
  rootMux.Handle("/", assets.SPAHandler("index.html"))

  http.ListenAndServe(":1323", rootMux)
}
