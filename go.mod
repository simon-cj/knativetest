module knativetest

go 1.15

require (
	github.com/docker/docker v1.13.1
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/tektoncd/pipeline v0.18.1
	go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738
	golang.org/x/tools v0.0.0-20201022035929-9cf592e881e9 // indirect
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/cluster-bootstrap v0.20.4 // indirect
	k8s.io/klog v1.0.0
	k8s.io/kube-proxy v0.20.4 // indirect
	k8s.io/kubelet v0.20.4 // indirect
	knative.dev/networking v0.0.0-20201103163404-b9f80f4537af
	knative.dev/serving v0.19.0
)

// Pin k8s deps to v0.18.8
replace (
	k8s.io/api => k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/apiserver => k8s.io/apiserver v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/code-generator => k8s.io/code-generator v0.18.8
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29
)
