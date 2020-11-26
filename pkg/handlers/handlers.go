package handlers

import (
	"github.com/golang/glog"
	"net/http"
)

func HandleHome(w http.ResponseWriter, req *http.Request) error {
	glog.Errorf("Home called %s", req.Method)
	return nil
}