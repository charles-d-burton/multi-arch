package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//APIReq example json structure
type APIReq struct {
	Hello string `json:"hello"`
}

func main() {
	log.Println("Starting simple API service")
	http.HandleFunc("/", apiRequest)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func apiRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Received, request")
	var apiReq APIReq
	apiReq.Hello = "world"
	json.NewEncoder(w).Encode(apiReq)
}
