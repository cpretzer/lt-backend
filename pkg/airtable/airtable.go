package airtable

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
)

const (
	airtableKeyVariable  = "AIRTABLE_KEY"
	airtableBaseVariable  = "AIRTABLE_BASE"
	airtableHostVariable = "AIRTABLE_HOST"
	defaultAirtableHost = "https://api.airtable.com/v0/"
)

type AirTableClientInterface interface{}

type AirtableClient struct {
	Key    *string
	Url    *string
	Client http.Client
}

type AirtableRequest struct {
	Method string
	Table string
	BytesReader io.Reader
}

type AirtableRecord struct {
	CreatedTime string     `json:"createdTime,omitempty"`
	Fields      interface{} `json:"fields,omitempty"`
	Id          string     `json:"id,omitempty"`
}

type AirtablePayload struct {
	Records []AirtableRecord `json:"records,omitempty"`
}

func InitializeClient() (*AirtableClient, error) {

	// for glog and anything else
	flag.Parse()

	glog.Info("Starting airtable service")

	airtableUrl, err := generateAirtableURL()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to generate AirTable URL %v", err))
	}

	airtableKey, isSet := os.LookupEnv(airtableKeyVariable)

	if !isSet || airtableKey == "" {
		return nil, errors.New("The AIRTABLE_KEY environment variable is not set")
	}

	return &AirtableClient{
		Key:    &airtableKey,
		Url:    airtableUrl,
		Client: initAirtableClient(),
	}, nil

}

func (c *AirtableClient) SendRequest(req *AirtableRequest) ([]byte, error) {
	url := fmt.Sprintf(*c.Url, req.Table)
	
	httpReq, err := req.buildHttpRequest(url, c.Key)
	if err != nil {
		glog.Errorf("Error sending the AirtableRequest %s", err)
		return nil, err
	}

	glog.Infof("Generated HTTP request %s", httpReq.Header.Get(authorizationHeader))
	glog.Infof("Sending request to %s using key %s", url, *c.Key)

	resp, err := c.Client.Do(httpReq)
	if err != nil {
		glog.Errorf("Error sending request to airtable %s", err)
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		glog.Errorf("There was an error reading the response body %s", err)
	}

	glog.V(8).Infof("Got response body %v", string(respBytes))

	defer resp.Body.Close()
	
	return respBytes, nil
}

func (r *AirtableRequest) buildHttpRequest(url string, key *string) (*http.Request, error) {
	httpReq, err := http.NewRequest(
		r.Method,
		url,
		r.BytesReader)

	if err != nil {
		glog.Errorf("There was an error building the HTTP request %s", err)
		return nil, err
	}

	httpReq.Header.Add(contentHeader, jsonUtf8)
	httpReq.Header.Add(authorizationHeader, fmt.Sprintf(bearerString, *key))

	return httpReq, nil
}


func generateAirtableURL() (*string, error) {
	airtableBaseId, isSet := os.LookupEnv(airtableBaseVariable)

	if !isSet {
		return nil, errors.New("AirTable App ID is not set")
	}

	airtableHost, isSet := os.LookupEnv(airtableHostVariable)

	if !isSet {
		airtableHost = defaultAirtableHost
	}

	airtableBaseUrl := fmt.Sprintf(airtableHost + "%s", airtableBaseId) + "/%s"

	// url := fmt.Sprintf(airtableBaseUrl, airtableBaseVariable)

	glog.Infof("Initialized Airtable URL: %v", airtableBaseUrl)

	return &airtableBaseUrl, nil
}

func initAirtableClient() http.Client {
	return http.Client{
		Timeout: time.Second * 15,
	}
}
