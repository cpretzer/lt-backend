package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"github.com/golang/glog"
	at "github.com/cpretzer/lt-backend/pkg/airtable"
)

func HandleGetUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	bytes, err := c.SendRequest(&at.AirtableRequest{
		Method: http.MethodGet,
		Table: "users",
		BytesReader: new(bytes.Buffer),
	})

	if err != nil {
		glog.Errorf("Error calling get users %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	glog.Infof("GetUser called %s", req.Method)
	return nil
}

func HandleAddUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	user := &User{
		FirstName: "Charles",
		LastName: "Pretzer",
		EmailAddress: "c@chabrina.com",
		Username: "cpretzer",
	}

	glog.V(8).Infof("Created user %+v", user)

	addUserRecord := &at.AirtableRecord{
		Fields: *user,
	}

	glog.V(8).Infof("Created addUserRecord %+v", addUserRecord)

	userRecords := make([]at.AirtableRecord, 0)

	userRecords = append(userRecords, *addUserRecord)

	glog.V(8).Infof("Created user records %+v", userRecords)

	addUserPayload := &at.AirtablePayload{
		Records: userRecords,
	}

	glog.V(8).Infof("Created add user payload %+v", addUserPayload)

	b := new(bytes.Buffer)

	err := json.NewEncoder(b).Encode(addUserPayload)

	if err != nil {
		glog.Errorf("Encoded Bytes error %s", err)
		return nil
	}
	

	bytes, err := c.SendRequest(&at.AirtableRequest{
		Method: http.MethodPost,
		Table: "users",
		BytesReader: b,
	})

	if err != nil {
		glog.Errorf("Error calling add user %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	glog.Infof("AddUser called %s", req.Method)
	return nil
}

func HandleDeleteUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {
	bytes, err := c.SendRequest(&at.AirtableRequest{
		Method: http.MethodGet,
		Table: "users",
		BytesReader: new(bytes.Buffer),
	})

	if err != nil {
		glog.Errorf("Error calling delete user %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	glog.Infof("DeleteUser called %s", req.Method)
	return nil
}

func HandleUpdateUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {
	bytes, err := c.SendRequest(&at.AirtableRequest{
		Method: http.MethodGet,
		Table: "users",
		BytesReader: new(bytes.Buffer),
	})

	if err != nil {
		glog.Errorf("Error calling update user %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	glog.Infof("UpdateUser called %s", req.Method)
	return nil
}