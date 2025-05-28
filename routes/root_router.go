package routes

import (
	"fmt"
	"main/assets"
	"net/http"
)

func RootRouter(port string) {
  mux := http.NewServeMux()
  mux.Handle("/api/", http.StripPrefix("/api", APIRouter()))
  mux.Handle("/", assets.SPAHandler("index.html"))
  fmt.Println("Listening on port", port)
  http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
