package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeRequest[T any](r *http.Request) (T, error) {
	var payload T

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func EncodeResponse[T any](w http.ResponseWriter, statusCode int, payload T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
