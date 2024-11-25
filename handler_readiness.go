package main

import "net/http"

// Http handle request, should respond is the server is ready and working
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
