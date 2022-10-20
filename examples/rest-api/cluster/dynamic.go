package cluster

import (
	"log"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func NewDynamicClusterClient() dynamic.Interface {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/wby/.kube/config")
	if err != nil {
		log.Fatalln("get cluster config error")
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalln("init cluster client error")
	}

	return dynamicClient
}
