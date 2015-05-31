package helpdesk

import (
	"github.com/fkse/helpdesk/controller"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var AppRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.DefaultIndex,
	},
	// User routes
	Route{
		"UserList",
		"GET",
		"/users",
		controller.UserList,
	},
	Route{
		"UserCreate",
		"POST",
		"/users",
		controller.UserCreate,
	},
	Route{
		"UserShow",
		"GET",
		"/users/{id:[0-9+]}",
		controller.UserShow,
	},
	Route{
		"UserUpdate",
		"PUT",
		"/users/{id:[0-9+]}",
		controller.UserUpdate,
	},
	Route{
		"UserDelete",
		"DELETE",
		"/users/{id:[0-9+]}",
		controller.UserDelete,
	},
}
