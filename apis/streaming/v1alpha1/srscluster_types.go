/*


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

type VHostSRSHLSConfigSpec struct {
	Enabled   bool    `json:"enabled,omitempty"`
	Fragment  *int32  `json:"hlsFragment,omitempty"`
	Window    *int32  `json:"hlsWindow,omitempty"`
	Path      *string `json:"hlsPath,omitempty"`
	M3U8File  *string `json:"hlsM3u8_file,omitempty"`
	TSFile    *string `json:"hlsTs_file,omitempty"`
	CleanUp   *bool   `json:"hlsCleanup,omitempty"`
	NBNotify  *int32  `json:"hlsNbNotify,omitempty"`
	WaitFrame *bool   `json:"hlsWaitKeyframe,omitempty"`
}

// VHostSRSConfigSpec
type VHostSRSConfigSpec struct {
	Name       string                 `json:"name"`
	Enabled    *bool                  `json:"enabled,omitempty"`
	MixCorrect *bool                  `json:"mixCorrect,omitempty"`
	HLS        *VHostSRSHLSConfigSpec `json:"hls,omitempty"`
}

type SRSConfigServerSpec struct {
	Enabled bool   `json:"enabled,omitempty"`
	Listen  *int32 `json:"listen,omitempty"`
}

// SRSConfigSpec Configuration
type SRSConfigSpec struct {
	Daemon         bool                 `json:"daemon,omitempty"`
	MaxConnextions *int32               `json:"maxConnections,omitempty"`
	Listen         *int32               `json:"listen,omitempty"`
	API            *SRSConfigServerSpec `json:"api,omitempty"`
	Server         *SRSConfigServerSpec `json:"server,omitempty"`
	VHosts         []VHostSRSConfigSpec `json:"vhosts,omitempty"`

	Origins []string `json:"-"`
}

// SRSClusterSpec defines the desired state of SRSCluster
type SRSClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Size ClusterOrigin size
	Size int32 `json:"size"`

	// Image version SRS, default: ossrs/srs:v4.0.56
	Image string `json:"image,omitempty"`

	// ServiceName Service name
	ServiceName string `json:"serviceName,omitempty"`

	// Config SRS configuration
	Config *SRSConfigSpec `json:"config,omitempty"`
}

// SRSClusterStatus defines the observed state of SRSCluster
type SRSClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	PodNames          []string `json:"podNames"`
	AvailableReplicas int32    `json:"availableReplicas"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SRSCluster is the Schema for the srsclusters API
type SRSCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SRSClusterSpec   `json:"spec,omitempty"`
	Status SRSClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SRSClusterList contains a list of SRSCluster
type SRSClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SRSCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SRSCluster{}, &SRSClusterList{})
}
