package goals

import (
	"encoding/json"
	"errors"
	"net/http"

	handlers "github.com/cpretzer/lt-backend/pkg/handlers"
	at "github.com/cpretzer/tavolo-dellaria"
	"github.com/golang/glog"
)

func HandleGetGoal(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	goalId := req.URL.Query().Get(handlers.IdParam)

	goal, err := doGetReport(c, goalId)

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

func HandleCreateReport(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	glog.V(5).Infof("create Report %v", req.Body)

	report, err := UnmarshalReport(req)

	if err != nil {
		glog.Errorf("Error unmarshaling the report request body %s", err)
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Unable to create goal",
		)
		return nil
	}

	addReportRequest := c.CreateAirtableRequest(http.MethodPost, reportsTable)

	glog.V(8).Infof("Created add report request %+v", addReportRequest)

	addReportRecord := addReportRequest.CreateRecord(report)

	addReportRequest.AddRecordToRequest(*addReportRecord)

	bytes, err := c.SendRequest(addReportRequest)

	if err != nil {
		glog.Errorf("Error calling add report %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))

	glog.Infof("CreateReport called %s", req.Method)
	return nil
}

// HandleUpdateReport  handles requests to the route defined in routes.go
// It takes a PUT request, transforms it, and sends it to Airtable
func HandleUpdateGoal(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	glog.V(8).Infof("HandleUpdateGoal %s", req.Body)

	// Get the report from the request body
	reportUpdate, err := UnmarshalReport(req)

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

	reportId := reportUpdate.ReportId

	existingReport, err := doGetReport(c, reportId)

	glog.V(8).Infof("got existing report %+v", existingReport)

	if existingReport == nil {
		glog.V(8).Infof("No report found")
		handlers.WriteError(w, &err,
			http.StatusBadRequest, "Unable to update report")
		return nil
	}

	existingReport.IsActive = reportUpdate.IsActive

	updateRequest := c.CreateAirtableRequest(http.MethodPatch, reportsTable)

	updateRecord := updateRequest.CreateRecord(existingReport)
	updateRecord.Id = existingReport.Id
	updateRequest.AddRecordToRequest(*updateRecord)

	bytes, err := c.SendRequest(updateRequest)

	if err != nil {
		glog.Errorf("Error calling update goal %s", err)
	}

	glog.Infof("Got bytes %s", string(bytes))

	glog.Infof("UpdateGoal called %s", req.Method)
	return nil
}

func getReportFromRequest(c *at.AirtableClient, req *http.Request) (*Report, error) {
	glog.V(8).Infof("GetReportFromRequest %s", req.Body)

	reportUpdate, err := UnmarshalReport(req)

	if err != nil {
		glog.Errorf("Error unmarshaling the report request body %s", err)
		return nil, errors.New("Unable to update report")
	}

	existingReport, err := doGetReport(c, reportUpdate.ReportId)

	if existingReport == nil {
		glog.V(8).Infof("No report found")
		return nil, errors.New("Unable to get report")
	}

	glog.V(8).Infof("Retrieved report %+v", existingReport)

	return existingReport, nil
}

func HandleDeleteReport(c *at.AirtableClient, w http.ResponseWriter, req *http.Request) error {

	report, err := getReportFromRequest(c, req)

	if err != nil {
		glog.Errorf("Error getting the report from the request %s", err)
		handlers.WriteError(w,
			&err,
			http.StatusBadRequest,
			"Error deleting the report")
		return nil
	}

	glog.V(8).Infof("Report active %+v", report.IsActive)

	if report.IsActive {
		deactivateRequest := c.CreateAirtableRequest(http.MethodPatch, reportsTable)
		report.IsActive = false
		// existingUser.DeactivationDate = time.Now().Unix()
		deactivateRecord := deactivateRequest.CreateRecord(report)
		deactivateRecord.Id = report.Id

		deactivateRequest.AddRecordToRequest(*deactivateRecord)

		bytes, err := c.SendRequest(deactivateRequest)

		if err != nil {
			glog.Errorf("Error calling delete report %s", err)
		}

		glog.Infof("Got bytes %s", string(bytes))

	} else {
		glog.V(8).Infof("Report is already deactivated")
	}

	return nil
}

func doGetReport(c *at.AirtableClient, id string) (*Report, error) {

	var err error
	if id == "" {
		glog.Errorf("HandleGetReport no ID parameter")
		err := errors.New("No report ID param")
		return nil, err
	}

	airtableReq := c.MakeGetRecordRequest(reportsTable, id)

	bytes, err := c.SendRequest(airtableReq)

	if err != nil {
		glog.Errorf("Error calling get report %s", err)
		return nil, err
	}

	glog.Infof("Got bytes %s", string(bytes))

	var atRecord at.AirtableRecord
	err = json.Unmarshal(bytes, &atRecord)

	if err != nil {
		glog.Errorf("Error unmarshaling the report record body %s", err)
		return nil, err
	}

	var report Report
	fieldMap := atRecord.Fields.(map[string]interface{})
	report.Id = fieldMap["id"].(string)

	if fieldMap["name"] != nil {
		report.Name = fieldMap["name"].(string)
	}

	// if fieldMap["description"] != nil {
	// 	report.Description = fieldMap["description"].(string)
	// }

	if fieldMap["isActive"] != nil {
		report.IsActive = fieldMap["isActive"].(bool)
	}

	if err != nil {
		glog.Errorf("Error unmarshaling the goal from the record %s", err)
		return nil, err
	}

	glog.V(8).Infof("returning report %+v", &report)

	return &report, err
}
