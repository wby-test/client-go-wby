package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/examples/rest-api/cluster"
)

func main() {
	fmt.Println("test")
	client := cluster.NewDynamicClient()
	//listDeployment(client)
	listPytorchJobs(client)
}

func listDeployment(p dynamic.Interface) {
	deployResource := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
	rt, err := p.Resource(deployResource).Namespace("public-resource").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, v := range rt.Items {
		fmt.Println(v.GetName())
	}

}

func listPytorchJobs(p dynamic.Interface) {
	torchResource := schema.GroupVersionResource{Group: "kubeflow.org", Version: "v1", Resource: "pytorchjobs"}
	rt, err := p.Resource(torchResource).Namespace("public-resource").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, v := range rt.Items {
		fmt.Println(v.GetName(), " ", v.GetCreationTimestamp(), v.UnstructuredContent())
	}

}
