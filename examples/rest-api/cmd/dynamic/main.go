package main

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/examples/rest-api/cluster"
)

func main() {
	dynamicClient := cluster.NewDynamicClusterClient()
	pytorchRes := schema.GroupVersionResource{
		Group:    "kubeflow.org",
		Version:  "v1",
		Resource: "PyTorchJob",
	}
	fmt.Println("-------------开始获取deployment-----------")
	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployment"}
	list, err := dynamicClient.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		replicas, found, err := unstructured.NestedInt64(d.Object, "spec", "replicas")
		if err != nil || !found {
			fmt.Printf("Replicas not found for deployment %s: error=%s", d.GetName(), err)
			continue
		}
		fmt.Printf(" * %s (%d replicas)\n", d.GetName(), replicas)
	}

	fmt.Println("------------开始获取pytorchjobs--------------")
	list, err = dynamicClient.Resource(pytorchRes).Namespace("public-resource").List(
		context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}
