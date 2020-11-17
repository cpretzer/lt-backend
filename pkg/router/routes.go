package main

import "net/http"

type Routes []Route

// Array of Route objects, each associated with a unique HTTP endpoint
var routes = Routes{
	// Handler listening for GET at "/" URI
	// Returns specified string
	Route{
		"Home",
		http.MethodGet,
		"/",
		HandleHome,
	},
}
