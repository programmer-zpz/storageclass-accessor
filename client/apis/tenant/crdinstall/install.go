package install

import (
	k8sruntime "k8s.io/apimachinery/pkg/runtime"
	urlruntime "k8s.io/apimachinery/pkg/util/runtime"
	tenantv1alpha1 "storageclass-accessor/client/apis/tenant/v1alpha1"
)

func Install(scheme *k8sruntime.Scheme) {
	urlruntime.Must(tenantv1alpha1.AddToScheme(scheme))
	urlruntime.Must(scheme.SetVersionPriority(tenantv1alpha1.SchemeGroupVersion))
}
