package cluster

import (
	"flag"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func NewDynamicClient(env string) dynamic.Interface {
	kubeConfig := GetConfigFile(env)
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return client
}
