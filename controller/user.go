package controller

import (
	"net/http"
	"github.com/fkse/helpdesk/model"
	"fmt"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	JsonError(w, "Internal server error", http.StatusInternalServerError)
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user model.User
	// get request bod
	//r.B
	fmt.Printf("%v", user)
}

func UserShow(w http.ResponseWriter, r *http.Request) {

}

func UserUpdate(w http.ResponseWriter, r *http.Request) {

}

func UserDelete(w http.ResponseWriter, r *http.Request) {

}
