// Package v1alpha1 contains API Schema definitions for the tenant v1alpha1 API group
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=cloudbases.io/cloudbases/pkg/apis/tenant
// +k8s:defaulter-gen=TypeMeta
// +groupName=tenant.cloudbases.io
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "tenant.cloudbases.io", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme is required by pkg/client/...
	AddToScheme = SchemeBuilder.AddToScheme
)

// Resource is required by pkg/client/listers/...
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
