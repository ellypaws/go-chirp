package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Encode[T any](w http.ResponseWriter, data T) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("error encoding response: %v", err)
	}
	return nil
}

func Decode[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("error decoding request body: %v", err)
	}
	return data, nil
}
