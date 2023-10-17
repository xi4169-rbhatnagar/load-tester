package server

import "net/http"

func registerRoutes(handler http.Handler) http.Handler {
	http.HandleFunc("/submit", HandleSubmit)
	return handler
}
