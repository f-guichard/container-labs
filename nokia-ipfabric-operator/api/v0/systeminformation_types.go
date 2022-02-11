/*
Copyright 2022.

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

package v0

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SystemInformationSpec defines the desired state of SystemInformation
type SystemInformationSpec struct {
	Infos InformationSpec 	`json:"information"`
	Name NameSpec 		`json:"name"`
	Dns DnsSpec 		`json:"dns"`
	Banner BannerSpec 	`json:"banner"`
	Ntp NtpSpec 		`json:"ntp"`
	Clock ClockSpec 	`json:"clock"`
	Snmp SnmpSpec 		`json:"snmp"`
	Logging LoggingSpec 	`json:"logging"`
}

type InformationSpec struct {
	Location string 	`json:"location"`
	Contact string		`json:"contact"`
}

type NameSpec struct {
	Hostname string		`json:"host-name"`
	DomainName string	`json:"domain-name"`
}

type DnsSpec struct {
	NetworkInstance string `json:"network-instance"`
	ServerList []string    `json:"server-list"`
	HostEntry []HostEntrySpec `json:"host-entry"`
}

// SystemInformationStatus defines the observed state of SystemInformation
type SystemInformationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SystemInformation is the Schema for the systeminformations API
type SystemInformation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemInformationSpec   `json:"spec,omitempty"`
	Status SystemInformationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SystemInformationList contains a list of SystemInformation
type SystemInformationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SystemInformation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SystemInformation{}, &SystemInformationList{})
}
