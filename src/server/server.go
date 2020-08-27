package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(res http.ResponseWriter, reg *http.Request) {
	res.Write([]byte("Hello from Go server"))
}
