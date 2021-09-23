package handlers

import (
	"net/http"

	at "github.com/cpretzer/lt-backend/pkg/airtable"
	"github.com/golang/glog"
)

func HandleHome(atClient *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	atReq := atClient.CreateAirtableRequest(http.MethodGet, "users")
	bytes, err := atClient.SendRequest(atReq)

	if err != nil {
		glog.Errorf("Error calling home %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))

	glog.Infof("Home called %s", req.Method)
	return nil
}
