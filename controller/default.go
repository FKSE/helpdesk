package controller

import (
	"encoding/json"
	"net/http"
)

func DefaultIndex(w http.ResponseWriter, r *http.Request) {
	sysinfo := map[string]string{
		"application": "fkse-helpdesk",
		"version":     "1.0.0",
	}
	// status
	w.WriteHeader(http.StatusOK)
	// content
	json.NewEncoder(w).Encode(sysinfo)
}
