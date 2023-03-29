module cloudbases.io/storageclass-accessor

go 1.16

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-openapi/spec v0.19.7
	github.com/spf13/cobra v1.2.1
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v12.0.0+incompatible // indirect
	k8s.io/klog/v2 v2.9.0
	k8s.io/kube-openapi v0.0.0-20210421082810-95288971da7e
	sigs.k8s.io/controller-runtime v0.10.0
)

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1
	k8s.io/client-go => k8s.io/client-go v0.21.2
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7
)
