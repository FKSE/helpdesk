package controller

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const (
	MaxRequestBody = 26214400
)

//

// TODO: Implement :D
func IsAuthorized(role string, r *http.Request) bool {
	return true
}

func DecodeRequestBody(w http.ResponseWriter, r *http.Request, v *interface{}) (err error) {
	body, err := ioutil.ReadAll(http.MaxBytesReader(w, r.Body, MaxRequestBody))
	if err != nil {
		// sig client
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return err
	}
	// decode json data
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}

type ApiError struct {
	Id string //Request url + time encoded
	Status string
	Code string
	Title string
	Detail string
}

func JsonError(w http.ResponseWriter, msg string, code int) {
	err := map[string]map[string]interface{}{
		"errors": {
			"title": msg,
			"status":code,
		},
	}
	// status
	w.WriteHeader(code)
	// content
	json.NewEncoder(w).Encode(err)
}