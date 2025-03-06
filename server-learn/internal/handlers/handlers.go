package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status":  "ok",
		"message": "Server is running",
	}

	json.NewEncoder(w).Encode(response)
}

// func DownloadHandler(w http.ResponseWriter, r *http.Request) {

// }
