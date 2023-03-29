package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindWorkspace     = "Workspace"
	ResourceSingularWorkspace = "workspace"
	ResourcePluralWorkspace   = "workspaces"
	WorkspaceLabel            = "cloudbases.io/workspace"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	Manager          string `json:"manager,omitempty"`
	NetworkIsolation *bool  `json:"networkIsolation,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace
type WorkspaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// Workspace is the Schema for the workspaces API
// +k8s:openapi-gen=true
// +kubebuilder:resource:categories="tenant",scope="Cluster"
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec,omitempty"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// WorkspaceList contains a list of Workspace
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
