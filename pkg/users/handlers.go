package users

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"github.com/golang/glog"
	at "github.com/cpretzer/lt-backend/pkg/airtable"
	handlers "github.com/cpretzer/lt-backend/pkg/handlers"
)

func HandleGetUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	userId := req.URL.Query().Get(handlers.IdParam)
	
	var err error
	if userId == "" {
		glog.Errorf("No ID parameter")
		err := errors.New("No ID param")
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to update user",
		)
		return nil
	}	
	
	getUserRequest := c.MakeGetRecordRequest(usersTable, userId)

	bytes, err := c.SendRequest(getUserRequest)

	if err != nil {
		glog.Errorf("Error calling get user %s", err)
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to update user",
		)
		return nil
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	var user User
	err = json.Unmarshal(bytes, &user)

	if err != nil {
		glog.Errorf("Error unmarshaling the user request body %s", err)
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to update user",
		)
		return nil
	}

	return nil
}

func HandleAddUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	glog.V(5).Infof("add user %v", req.Body)

	user, err := UnmarshalUser(req)

	if err != nil {
		glog.Errorf("Error unmarshaling the user request body %s", err)
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to update user",
		)
		return nil
	}

	addUserRequest := c.CreateAirtableRequest(http.MethodPost, usersTable)

	glog.V(8).Infof("Created add user request %+v", addUserRequest)

	// addUserRecord := &at.AirtableRecord{
	// 	Fields: *user,
	// }

	addUserRecord := addUserRequest.CreateRecord(user)

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

	glog.V(8).Infof("HandleDeleteUser %+v", req.Body)

	deleteUser, err := UnmarshalUser(req)

	if err != nil {
		glog.Errorf("Error Unmarshaling user request body %s", err)
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to deactivate account")
	}

	existingUser, err := RetrieveUserByEmailAddress(deleteUser.EmailAddress, c)

	if existingUser == nil {
		glog.V(8).Infof("No user found for email %s", deleteUser.EmailAddress)	
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to deactivate account")
	}

	if existingUser.Active {
		deactivateRequest := c.CreateAirtableRequest(http.MethodPatch, usersTable)
		existingUser.Active = false
		existingUser.DeactivationDate = time.Now().Unix()
		deactivateRecord := deactivateRequest.CreateRecord(existingUser)
		deactivateRecord.Id = existingUser.Id
		
		deactivateRequest.AddRecordToRequest(*deactivateRecord)

		bytes, err := c.SendRequest(deactivateRequest)

		if err != nil {
			glog.Errorf("Error calling delete user %s", err)
		}
	
		glog.Infof("Got bytes %s", string(bytes))
				
	} else {
		glog.V(8).Infof("User is already deactivated")
	}

	
	glog.Infof("DeleteUser called %s", req.Method)
	return nil
}

func HandleUpdateUser(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	glog.V(8).Infof("HandleUpdateUser %s", req.Body)

	userUpdate, err := UnmarshalUser(req)

	if err != nil {
		glog.Errorf("Error unmarshaling the user request body %s", err)
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to update user",
		)
		return nil
	}

	existingUser, err := RetrieveUserByEmailAddress(userUpdate.EmailAddress, c)

	if existingUser == nil {
		glog.V(8).Infof("No user found")
		handlers.WriteError(w, &err, 
			http.StatusBadRequest, "Unable to update user")
		return nil
	}

	existingUser.FirstName = userUpdate.FirstName
	existingUser.LastName = userUpdate.LastName
	existingUser.Username = userUpdate.Username

	updateRequest := c.CreateAirtableRequest(http.MethodPatch, usersTable)

	updateRecord := updateRequest.CreateRecord(existingUser)
	updateRecord.Id = existingUser.Id
	updateRequest.AddRecordToRequest(*updateRecord)

	bytes, err := c.SendRequest(updateRequest)

	if err != nil {
		glog.Errorf("Error calling update user %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	glog.Infof("UpdateUser called %s", req.Method)
	return nil
}

func getUserFromRequest(c *at.AirtableClient, req *http.Request) (*User, error) {
	glog.V(8).Infof("HandleUpdateUser %s", req.Body)

	userUpdate, err := UnmarshalUser(req)

	if err != nil {
		glog.Errorf("Error unmarshaling the user request body %s", err)
		return nil, errors.New("Unable to update user")
	}

	existingUser, err := RetrieveUserByEmailAddress(userUpdate.EmailAddress, c)

	if existingUser == nil {
		glog.V(8).Infof("No user found")
		return nil, errors.New("Unable to update user")
	}

	return existingUser, nil
}