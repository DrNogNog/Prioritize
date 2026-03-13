package api

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/api/chat", ChatHandler)
	http.HandleFunc("/api/hello", HelloHandler)
}