/*
Copyright 2023.

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

// MetadataSpec defines the desired state of Metadata
type MetadataSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Metadata. Edit metadata_types.go to remove/update
	Team    string `json:"team,omitempty"`
	Contact string `json:"contact,omit"`
}

// MetadataStatus defines the observed state of Metadata
type MetadataStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Information when was the last time the job was successfully scheduled.
	// +optional
	// LastSuccessTime *metav1.Time `json:"lastSuccessTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Metadata is the Schema for the metadata API
type Metadata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MetadataSpec   `json:"spec,omitempty"`
	Status MetadataStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MetadataList contains a list of Metadata
type MetadataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Metadata `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Metadata{}, &MetadataList{})
}
