package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// logMiddleware implements pattern Decorator
func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from logMiddleware")
		next(w, r)
	}
}

// handleHome is a simple handler
func handleHome(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello from handleHome")
	w.WriteHeader(http.StatusOK)
}

func startServer() {
	// http.HandleFunc("/", handleHome)
	http.HandleFunc("/", logMiddleware(handleHome))
	http.ListenAndServe(":8080", nil)
}

func main() {
	startServer()
}
