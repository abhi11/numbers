package util

import "net/http"

// Util method to write the response back
func WriteResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
