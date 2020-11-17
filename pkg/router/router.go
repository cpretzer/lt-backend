package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Router for associating HTTP requests with functions based on URI
// Router takes parameters: Name, Method, Path, and Handler to associate
// with a function
func NewRouter() *mux.Router {

	// Create new gorilla/mux router with with strict slash
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		// Associate each route with an HTTP endpoint
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(Handler{env, route.Function})

	}

	// Return router to be used by server
	return router
}

type HttpResponseError struct {
	errorMessage string
}

func NewHttpResponseError(msg string) error {
	return &HttpResponseError{
		errorMessage: msg,
	}
}

func (e *HttpResponseError) Error() string {
	return e.errorMessage
}

