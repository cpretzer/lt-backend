package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	at "github.com/cpretzer/lt-backend/pkg/airtable"
)

func UnmarshalUser(req *http.Request) (*User, error) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		glog.Errorf("Unable to read update user body %s", err)
		return nil, err
	}

	var user User

	err = json.Unmarshal(body, &user)

	if err != nil {
		glog.Errorf("Unable to unmarshal body %s, [%s]", string(body), err)
		return nil, err
	}

	glog.V(8).Infof("Unmarshaled user %+v", user)

	return &user, nil
}

func RetrieveUserByEmailAddress(emailAddress string, c *at.AirtableClient) (*User, error) {

	queryString := fmt.Sprintf("%v%v", at.FilterQueryString, 
		fmt.Sprintf(userEmailFilter, emailAddress))

	glog.V(5).Infof("airtable query string %v", queryString)

	userEmailRequest := c.MakeFilterRecordRequest(usersTable, queryString)
	

	b, err := c.SendRequest(userEmailRequest)

	if err != nil {
		glog.Errorf("Received error when querying airtable %v", err)
		return nil, errors.New("received error when querying airtable for user by email")
	}

	glog.V(8).Infof("bytes response %v", string(b))

	var resp at.AirtablePayload
	err = json.Unmarshal(b, &resp)

	if err != nil {
		glog.Errorf("Unable to unmarshal airtable payload %v", err)
		return nil, errors.New("unable to unmarshal airtable payload")
	}

	glog.V(5).Infof("airtable response has %d records", len(resp.Records))
	glog.V(5).Infof("airtable response: %v", resp)

	if len(resp.Records) < 1 {
		glog.V(8).Infof("There are no users with emailAddress %s", emailAddress)
		return nil, errors.New(
			fmt.Sprintf("No user found with this email address %s", emailAddress))
	}

	// get the first one for now

	var user User
	fieldMap := resp.Records[0].Fields.(map[string]interface{})
	user.Id = resp.Records[0].Id
	user.FirstName = fieldMap["firstName"].(string)
	user.LastName = fieldMap["lastName"].(string)
	user.EmailAddress = fieldMap["emailAddress"].(string)
	user.Username = fieldMap["username"].(string)

	if fieldMap["active"] != nil {		
		user.Active = fieldMap["active"].(bool)
	}

	if fieldMap["creationDate"] != nil {		
		user.CreationDate = fieldMap["creationDate"].(uint)
	}

	if fieldMap["lastLogin"] != nil {
		user.LastLogin = fieldMap["lastLogin"].(uint)
	}

	return &user, nil

}