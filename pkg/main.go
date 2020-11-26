/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (

	"flag"
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"os"
	"time"
	"github.com/cpretzer/lt-backend/pkg/router"
)

func main() {

	// for glog and anything else
	flag.Parse()

	glog.Info("GLOG: Starting invitations service")


	router := router.NewRouter()

	port, isSet := os.LookupEnv("SERVICE_PORT")

	if !isSet || port == "" {
		port = "3001"
	}

	glog.Info(fmt.Sprintf("Using service port %v", port))

	// Listen for requests on port :8080 with router and logging
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		glog.Fatalf("ListenAndServe error %s", err)
	}
}

func initAirTableClient() http.Client {
	return http.Client{
		Timeout: time.Second * 15,
	}
}
