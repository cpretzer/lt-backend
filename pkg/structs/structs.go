package structs

import (
	"net/http"
)

type Questionnaire struct {
	Questions []Question `json:"questions,omitempty"`
	User User `json:"question,omitempty"`
	CreationDate string `json:"creationDate,omitempty"`
}

type Question struct {
	Text string `json:"text,omitempty"`
	Category QuestionCategory `json:"questionCategory,omitempty"`
}

type QuestionCategory struct {
	Name string `json:"name,omitempty"`
}

type QuestionnaireResponse struct {
	Question Question `json:"question,omitempty"`
	CreationDate string `json:"creationDate,omitempty"`
}

type Report struct {
	Responses []QuestionnaireResponse `json:"responses,omitempty"`
	ReportDate string `json:"reportDate,omitemptry"`
}

type ReportSummary struct {
	SummaryDate string `json:"summaryDate"`
}

type User struct {
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Email string `json:"email,omitempty"`
	LastLogin string `json:"lastLogin,omitempty"`
	Password string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
}

type Route struct {
	Name     string
	Method   string
	Pattern  string
	Function HandlerFunc
}

// HandlerFunc type used to specify a template for handler functions to follow
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// Handler object used for allowing handler functions to accept
// an environment object
type Handler struct {
	H func(w http.ResponseWriter, r *http.Request) error
}

type Error interface {
	error
	Status() int
}

