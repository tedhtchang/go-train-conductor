/*
Copyright 2024 Train Conductor Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TuningJobSpec defines the desired state of TuningJob
type TuningJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of TuningJob. Edit tuningjob_types.go to remove/update
	ModelName  string             `json:"ModelName"`
	OutputPath string             `json:"OutputPath"`
	Parameters TrainingParameters `json:"Parameters"`
}
type TrainingParameters struct {
	ModelNameOrPath string `json:"ModelNameOrPath"`
	DataPath        string `json:"DataPath"`
	SaveStrategy    string `json:"SaveStrategy,omitempty"`
}

// TuningJobStatus defines the observed state of TuningJob
type TuningJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State TrainingStatus `json:"TrainingStatus"`
}

type TrainingStatus string

const (
	Placeholderunset TrainingStatus = "placeholderunset"
	Pending          TrainingStatus = "pending"
	Queued           TrainingStatus = "queued"
	Running          TrainingStatus = "running"
	Suspended        TrainingStatus = "suspended"
	Completed        TrainingStatus = "completed"
	Canceled         TrainingStatus = "canceled"
	Failed           TrainingStatus = "failed"
	Deleted          TrainingStatus = "deleted"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TuningJob is the Schema for the tuningjobs API
type TuningJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TuningJobSpec   `json:"spec,omitempty"`
	Status TuningJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TuningJobList contains a list of TuningJob
type TuningJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TuningJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TuningJob{}, &TuningJobList{})
}
