package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// native http

	/*
		// creates new goroutine from handler for each request
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
	*/

	// mux

	router := mux.NewRouter()
	router.HandleFunc("/", GetHomeController).Methods(http.MethodGet)
	router.HandleFunc("/", PostHomeController).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}

// LoginRequest is a body of http request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetHomeController - GET /
func GetHomeController(rw http.ResponseWriter, req *http.Request) {
	// extract query params from url ?key=value
	// mux.Vars(req)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	// rw.Write([]byte("Hello from Go server"))
	io.WriteString(rw, `{"alive": true}`) // instead of rw.Write
}

// PostHomeController - POST /
func PostHomeController(rw http.ResponseWriter, req *http.Request) {
	loginRequest := LoginRequest{}
	err := json.NewDecoder(req.Body).Decode(&loginRequest)

	if err != nil {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		io.WriteString(rw, `{"error":"`+err.Error()+`"}`)
		return
	}

	if len(loginRequest.Username) < 1 || len(loginRequest.Password) < 1 {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		io.WriteString(rw, `{"error":"Invalid data"}`)
		return
	}

	log.Println("PostHomeController loginRequest", loginRequest)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(loginRequest)
}
