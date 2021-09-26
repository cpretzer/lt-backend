package router

import (
	"github.com/cpretzer/lt-backend/pkg/structs"
	at "github.com/cpretzer/tavolo-dellaria"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"net/http"
)

// Router for associating HTTP requests with functions based on URI
// Router takes parameters: Name, Method, Path, and Handler to associate
// with a function
func NewRouter(client *at.AirtableClient) *mux.Router {

	// Create new gorilla/mux router with with strict slash
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		// Associate each route with an HTTP endpoint
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(Handler{client, route.Function})

	}

	// Return router to be used by server
	return router
}

type HttpResponseError struct {
	errorMessage string
}

// Handler object used for allowing handler functions to accept
// an environment object
type Handler struct {
	AirtableClient *at.AirtableClient
	H              func(atClient *at.AirtableClient, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP is called on each HTTP request. Specifies which function is
// called as well as how errors are handled and how logging is set
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.AirtableClient, w, r)
	if err != nil {
		switch e := err.(type) {
		case structs.Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			glog.Infof("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

func NewHttpResponseError(msg string) error {
	return &HttpResponseError{
		errorMessage: msg,
	}
}

func (e *HttpResponseError) Error() string {
	return e.errorMessage
}
