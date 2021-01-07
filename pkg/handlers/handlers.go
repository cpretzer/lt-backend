package handlers

import (
	"github.com/golang/glog"
	"net/http"
	at "github.com/cpretzer/lt-backend/pkg/airtable"
)

func HandleHome(atClient *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	bytes, err := atClient.SendRequest(&at.AirtableRequest{
		Method: http.MethodGet,
		Table: "users",
		Payload: nil,
	})

	if err != nil {
		glog.Errorf("Error calling home %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	glog.Infof("Home called %s", req.Method)
	return nil
}
