// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"cloud.google.com/go/logging"
	"fmt"
	"net/http"
	"time"
)

func (a *App) HelpHandler(w http.ResponseWriter, r *http.Request) {
	a.log.Log(logging.Entry{
		Severity: logging.Info,
		HTTPRequest: &logging.HTTPRequest{
			Request: r,
		},
		Labels:  map[string]string{"arbitraryField": "custom entry"},
		Payload: "Structured logging example.",
	})

	startTime := time.Now()
	time.Sleep(time.Second * 3)
	endTime := time.Now()
	layout := "2006-01-02 15:04:05.999"
	_, _ = fmt.Fprintf(w, fmt.Sprintf("start: %s; end: %s\n", startTime.Format(layout), endTime.Format(layout)))
	//	fmt.Fprintf(w, "Hello World!\n")
}
