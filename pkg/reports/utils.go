package goals

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

func UnmarshalReport(req *http.Request) (*Report, error) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		glog.Errorf("Unable to read report request body %s", err)
		return nil, err
	}

	var report Report

	// need to figure out how to manage the ID
	err = json.Unmarshal(body, &report)

	if err != nil {
		glog.Errorf("Unable to unmarshal body %s, [%s]", string(body), err)
		return nil, err
	}

	glog.V(8).Infof("Unmarshaled report %+v", report)

	if report.ReportId == "" {
		glog.Errorf("The report is nil")
		return nil, fmt.Errorf("can't find the report in %s", string(body))
	}

	return &report, nil
}
