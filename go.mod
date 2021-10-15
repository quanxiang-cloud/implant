module github.com/quanxiang-cloud/implant

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/quanxiang-cloud/overseer v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210928044308-7d9f5e0b762b // indirect
	golang.org/x/sys v0.0.0-20210917161153-d61c044b1678 // indirect
	google.golang.org/grpc v1.40.0
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	k8s.io/klog v1.0.0
	sigs.k8s.io/controller-runtime v0.10.1
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace github.com/quanxiang-cloud/overseer => ../../quanxiang-cloud/overseer
