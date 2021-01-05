package users

import (
	"net/http"
	"github.com/golang/glog"
	at "github.com/cpretzer/lt-backend/pkg/airtable"
)

func HandleGetUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {
	glog.Errorf("GetUsers called %s", req.Method)
	return nil
}

func HandleAddUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {
	glog.Errorf("GetUsers called %s", req.Method)
	return nil
}

func HandleDeleteUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {
	glog.Errorf("GetUsers called %s", req.Method)
	return nil
}

func HandleUpdateUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {
	glog.Errorf("GetUsers called %s", req.Method)
	return nil
}