package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	JsonResponse(w, status, map[string]string{"error": message})
}
