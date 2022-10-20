package cluster

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClusterClient() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/wby/.kube/config")
	if err != nil {
		log.Fatalln("get cluster config error")
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("init cluster client error")
	}

	return clientSet
}
