package cluster

import (
	"fmt"

	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

const (
	PathBase    = ".kube/kube.config"
	FileTest    = "kubeflow.test"
	FileProduct = "kubeflow.prod"
	FileDevelop = "kubeflow.dev"
)

func GetConfigFile(env string) string {
	homeDir := homedir.HomeDir()
	var config string
	fmt.Println(homeDir)
	switch env {
	case "test":
		config = FileTest
	case "product":
		config = FileProduct
	case "develop":
		config = FileDevelop
	default:
		klog.Fatalf("get k8s config error")
	}

	return homeDir + "/" + PathBase + "/" + config
}
