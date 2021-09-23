package goals

import (
	"encoding/json"
	"errors"
	"net/http"
	"github.com/golang/glog"
	at "github.com/cpretzer/lt-backend/pkg/airtable"
	handlers "github.com/cpretzer/lt-backend/pkg/handlers"
)

func HandleGetGoal(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {
	
	goalId := req.URL.Query().Get(handlers.IdParam)

	goal, err := doGetGoal(c, goalId)

	if err != nil {
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to get goal",
		)
		return nil
	}

	glog.V(8).Infof("Got goal %+v", goal)

	handlers.WriteResponse(w,
		goal,
		http.StatusOK)

	return nil
}

func HandleCreateGoal(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	glog.V(5).Infof("create goal %v", req.Body)

	goal, err := UnmarshalGoal(req)

	if err != nil {
		glog.Errorf("Error unmarshaling the goal request body %s", err)
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to create goal",
		)
		return nil
	}

	addGoalRequest := c.CreateAirtableRequest(http.MethodPost, goalsTable)

	glog.V(8).Infof("Created add goal request %+v", addGoalRequest)

	addGoalRecord := addGoalRequest.CreateRecord(goal)

	addGoalRequest.AddRecordToRequest(*addGoalRecord)	

	bytes, err := c.SendRequest(addGoalRequest)

	if err != nil {
		glog.Errorf("Error calling add goal %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	glog.Infof("CreateGoal called %s", req.Method)
	return nil
}

// HandleUpdateGoal handles requests to the route defined in routes.go
// It takes a PUT request, transforms it, and sends it to Airtable
func HandleUpdateGoal(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	glog.V(8).Infof("HandleUpdateGoal %s", req.Body)

	// Get the goal from the request body
	goalUpdate, err := UnmarshalGoal(req)

	// If there is an error, propagate it back to the client
	if err != nil {
		glog.Errorf("Error unmarshaling the goal request body %s", err)
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to update goal",
		)
		return nil
	}

	goalId := goalUpdate.GoalId

	existingGoal, err := doGetGoal(c, goalId)

	glog.V(8).Infof("got existing goal %+v", existingGoal)

	if existingGoal == nil {
		glog.V(8).Infof("No goal found")
		handlers.WriteError(w, &err, 
			http.StatusBadRequest, "Unable to update goal")
		return nil
	}

	existingGoal.IsSystem = goalUpdate.IsSystem
	existingGoal.IsActive = goalUpdate.IsActive

	updateRequest := c.CreateAirtableRequest(http.MethodPatch, goalsTable)

	updateRecord := updateRequest.CreateRecord(existingGoal)
	updateRecord.Id = existingGoal.Id
	updateRequest.AddRecordToRequest(*updateRecord)

	bytes, err := c.SendRequest(updateRequest)

	if err != nil {
		glog.Errorf("Error calling update goal %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	glog.Infof("UpdateGoal called %s", req.Method)
	return nil
}

func getGoalFromRequest(c *at.AirtableClient, req *http.Request) (*Goal, error) {
	glog.V(8).Infof("GetGoalFromRequest %s", req.Body)

	goalUpdate, err := UnmarshalGoal(req)

	if err != nil {
		glog.Errorf("Error unmarshaling the goal request body %s", err)
		return nil, errors.New("Unable to update goal")
	}

	existingGoal, err := doGetGoal(c, goalUpdate.GoalId)

	if existingGoal == nil {
		glog.V(8).Infof("No goal found")
		return nil, errors.New("Unable to get goal")
	}

	glog.V(8).Infof("Retrieved goal %+v", existingGoal)

	return existingGoal, nil
}

func HandleDeleteGoal(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	goal, err := getGoalFromRequest(c, req)

	if err != nil {
		glog.Errorf("Error getting the goal from the request %s", err)
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Error deleting the goal")
		return nil
	}

	glog.V(8).Infof("Goal active %+v", goal.IsActive)

	if goal.IsActive {
		deactivateRequest := c.CreateAirtableRequest(http.MethodPatch, goalsTable)
		goal.IsActive = false
		// existingUser.DeactivationDate = time.Now().Unix()
		deactivateRecord := deactivateRequest.CreateRecord(goal)
		deactivateRecord.Id = goal.Id
		
		deactivateRequest.AddRecordToRequest(*deactivateRecord)

		bytes, err := c.SendRequest(deactivateRequest)

		if err != nil {
			glog.Errorf("Error calling delete user %s", err)
		}
	
		glog.Infof("Got bytes %s", string(bytes))
				
	} else {
		glog.V(8).Infof("User is already deactivated")
	}

	return nil
}

func doGetGoal(c *at.AirtableClient, id string) (*Goal, error) {
	
	var err error
	if id == "" {
		glog.Errorf("HandleGetGoal no ID parameter")
		err := errors.New("No goal ID param")
		return nil, err
	}	
	
	airtableReq := c.MakeGetRecordRequest(goalsTable, id)

	bytes, err := c.SendRequest(airtableReq)

	if err != nil {
		glog.Errorf("Error calling get goal %s", err)
		return nil, err
	}

	glog.Infof("Got bytes %s", string(bytes))
	
	var atRecord at.AirtableRecord
	err = json.Unmarshal(bytes, &atRecord)

	if err != nil {
		glog.Errorf("Error unmarshaling the goal record body %s", err)
		return nil, err
	}	

	var goal Goal
	fieldMap := atRecord.Fields.(map[string]interface{})
	goal.Id = fieldMap["id"].(string)

	if fieldMap["category"] != nil {
		goal.Category = fieldMap["category"].(string)
	}

	if fieldMap["name"] != nil {
		goal.Name = fieldMap["name"].(string)
	}

	if fieldMap["description"] != nil {
		goal.Description = fieldMap["description"].(string)
	}
	
	if fieldMap["isActive"] != nil {
		goal.IsActive = fieldMap["isActive"].(bool)
	}
	
	if fieldMap["isSystem"] != nil {
		goal.IsActive = fieldMap["isSystem"].(bool)
	}

	if err != nil {
		glog.Errorf("Error unmarshaling the goal from the record %s", err)
		return nil, err
	}	

	glog.V(8).Infof("returning goal %+v", &goal)

	return &goal, err
}