package install

import (
	tenantv1alpha1 "cloudbases.io/storageclass-accessor/client/apis/tenant/v1alpha1"
	k8sruntime "k8s.io/apimachinery/pkg/runtime"
	urlruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func Install(scheme *k8sruntime.Scheme) {
	urlruntime.Must(tenantv1alpha1.AddToScheme(scheme))
	urlruntime.Must(scheme.SetVersionPriority(tenantv1alpha1.SchemeGroupVersion))
}
