// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

/*
 * Airflow API (Stable)
 *
 * Apache Airflow management API.
 *
 * API version: 1.0.0
 * Contact: dev@airflow.apache.org
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package airflow
import (
	"time"
)
// DagDetail struct for DagDetail
type DagDetail struct {
	DagId string `json:"dag_id,omitempty"`
	RootDagId *string `json:"root_dag_id,omitempty"`
	IsPaused *bool `json:"is_paused,omitempty"`
	IsSubdag bool `json:"is_subdag,omitempty"`
	Fileloc string `json:"fileloc,omitempty"`
	// The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change. 
	FileToken string `json:"file_token,omitempty"`
	Owners []string `json:"owners,omitempty"`
	Description *string `json:"description,omitempty"`
	ScheduleInterval ScheduleInterval `json:"schedule_interval,omitempty"`
	Tags *[]Tag `json:"tags,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Catchup bool `json:"catchup,omitempty"`
	Orientation string `json:"orientation,omitempty"`
	Concurrency float32 `json:"concurrency,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	DagRunTimeout TimeDelta `json:"dag_run_timeout,omitempty"`
	DocMd string `json:"doc_md,omitempty"`
	DefaultView string `json:"default_view,omitempty"`
}
