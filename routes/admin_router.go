package routes

import (
	"main/middleware"
	"net/http"
)


func AdminRouter() *http.ServeMux {
	stack := middleware.MiddlewareStack(
		middleware.AdminMiddleware,
	)
	mux := http.NewServeMux()
	mux.Handle("/profile", stack(http.HandlerFunc(AdminHandler)))
	mux.Handle("/dashboard", stack(http.HandlerFunc(AdminDashboardHandler)))
	return mux
}