package cluster

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClusterClient() *kubernetes.Clientset {

	config, err := clientcmd.BuildConfigFromFlags("", "/Users/wangbaoyi1/.kube/config")
	if err != nil {
		panic(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return client
}
