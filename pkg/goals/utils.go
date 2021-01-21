package goals

import (
	"encoding/json"
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

	return &goal, nil
}