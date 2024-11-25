package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	//error codes 500 and above means bugs on my end
	if code > 499 {
		log.Println("Respodning with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

// takes a response writer and the status code and takes the interface
// Will marshal the payload into a JSON string
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	//adds a response header to the http request
	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
