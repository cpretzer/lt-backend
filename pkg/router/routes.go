package router

import (
	"net/http"
	"github.com/cpretzer/lt-backend/pkg/structs"
	"github.com/cpretzer/lt-backend/pkg/handlers"
)

type Routes []structs.Route

// Array of Route objects, each associated with a unique HTTP endpoint
var routes = Routes{
	// Handler listening for GET at "/" URI
	// Returns specified string
	structs.Route{
		Name: "Home",
		Method: http.MethodGet,
		Pattern: "/",
		Function: handlers.HandleHome,
	},
}
