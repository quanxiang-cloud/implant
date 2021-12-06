module github.com/quanxiang-cloud/implant

go 1.16

require (
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/quanxiang-cloud/overseer v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.41.0
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	k8s.io/klog v1.0.0
	sigs.k8s.io/controller-runtime v0.10.1
)

replace (
	github.com/quanxiang-cloud/overseer => ../../quanxiang-cloud/overseer
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7
)
