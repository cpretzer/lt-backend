package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/golang/glog"
	at "github.com/cpretzer/lt-backend/pkg/airtable"
	handlers "github.com/cpretzer/lt-backend/pkg/handlers"
)

func HandleGetUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	userId := req.URL.Query().Get(handlers.IdParam)

	if userId == "" {
		glog.Errorf("No ID parameter")
		handlers.WriteResponse(w, &handlers.JsonResponse{
			Code: http.StatusBadRequest,
			Message: "No user ID param", 
		}, http.StatusBadRequest)
	}	
	
	getUserRequest := c.MakeGetRecordRequest(usersTable, userId)

	bytes, err := c.SendRequest(getUserRequest)

	if err != nil {
		glog.Errorf("Error calling get user %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	var user User
	err = json.Unmarshal(bytes, &user)

	if err != nil {
		glog.Errorf("Error unmarshaling the user request body %s", err)
	}

	return nil
}

func HandleAddUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	glog.V(5).Infof("add user %v", req.Body)
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		glog.Errorf("unable to parse body %v", err)
		handlers.WriteResponse(w, &handlers.JsonResponse{
			Code:    http.StatusBadRequest,
			Message: "[TBDXXX] Unable to update add user",
		}, http.StatusBadRequest)
		return nil
	}

	var addUser User
	err = json.Unmarshal(body, &addUser)

	if err != nil {
		glog.Errorf("Unable to unmarshal body %s, [%s]", string(body), err)
	}

	glog.V(8).Infof("Created user %+v", addUser)

	addUserRequest := c.CreateAirtableRequest(http.MethodPost, usersTable)

	glog.V(8).Infof("Created add user request %+v", addUserRequest)

	// addUserRecord := &at.AirtableRecord{
	// 	Fields: *user,
	// }

	addUserRecord := addUserRequest.CreateRecord(addUser)

	addUserRequest.AddRecordToRequest(*addUserRecord)	

	bytes, err := c.SendRequest(addUserRequest)

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
		Payload: nil,
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
		Payload: nil,
	})

	if err != nil {
		glog.Errorf("Error calling update user %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	glog.Infof("UpdateUser called %s", req.Method)
	return nil
}