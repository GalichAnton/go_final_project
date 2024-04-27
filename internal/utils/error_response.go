package utils

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, error string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": error})
}
