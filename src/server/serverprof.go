package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// LogMiddleware implements pattern Decorator
func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from logMiddleware")
		next(w, r)
	}
}

// HandleHome is a simple handler
func HandleHome(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello from handleHome")
	w.WriteHeader(http.StatusOK)
}

func startServer() {
	// http.HandleFunc("/", HandleHome)
	http.HandleFunc("/", LogMiddleware(HandleHome))
	http.ListenAndServe(":8080", nil)
}

func main() {
	startServer()
}
