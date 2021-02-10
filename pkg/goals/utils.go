package goals

import (
	"fmt"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"github.com/golang/glog"	
)

func UnmarshalGoal(req *http.Request) (*Goal, error) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		glog.Errorf("Unable to read update user body %s", err)
		return nil, err
	}

	var goal Goal

	err = json.Unmarshal(body, &goal)

	if err != nil {
		glog.Errorf("Unable to unmarshal body %s, [%s]", string(body), err)
		return nil, err
	}

	glog.V(8).Infof("Unmarshaled goal %+v", goal)

	if &goal == nil || goal.GoalId == "" {
		glog.Errorf("The goal is nil or there was no ID in the request body")
		return nil, errors.New(fmt.Sprintf("Can't find the goal in %s", string(body)))
	}

	return &goal, nil
}