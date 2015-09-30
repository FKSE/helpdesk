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

// As defined by http://jsonapi.org/format/#errors
type ApiError struct {
	Status int
	Code string
	Title string
	Detail string
	Meta map[string]interface{}
}

func JsonError(w http.ResponseWriter, err *ApiError) {
	r := map[string]*ApiError{
		"errors": err,
	}
	// status
	w.WriteHeader(err.Status)
	// content
	json.NewEncoder(w).Encode(r)
}