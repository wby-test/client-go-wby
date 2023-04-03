package main

import (
	"context"
	"fmt"
	"os"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/examples/rest-api/cluster"
	"k8s.io/client-go/examples/rest-api/log"
	"k8s.io/klog/v2"
)

func main() {
	dynamicClient := cluster.NewDynamicClient("test")
	deployExample(dynamicClient)

	pytorchJobs(dynamicClient)
	//ListApiResource(dynamicClient)
}

func deployExample(client dynamic.Interface) {
	fmt.Println("-------------开始获取deployment-----------")
	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
	list, err := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
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
}

func pytorchJobs(client dynamic.Interface) {
	pwd, err := os.Getwd()
	if err != nil {
		klog.Error(err.Error())
	}
	fmt.Println(pwd)
	log.InitKlog(pwd, time.Now().String())
	// make sure we flush before exiting
	defer klog.Flush()

	pytorchRes := schema.GroupVersionResource{
		Group:    "kubeflow.org",
		Version:  "v1",
		Resource: "pytorchjobs",
	}
	fmt.Println("------------开始获取pytorchjobs--------------")

	list, err := client.Resource(pytorchRes).Namespace("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, v := range list.Items {
		fmt.Println(v.GetNamespace(), v.GetName())
		klog.Info(v.GetName(), v.GetNamespace(), v.UnstructuredContent())
		conditions, ok, err := unstructured.NestedSlice(v.Object, "status", "conditions")
		if err != nil {
			fmt.Println(conditions, ok)
		}
		status := conditions[len(conditions)-1].(map[string]interface{})["type"].(string)
		if status == "Failed" && v.GetCreationTimestamp().Time.Before(time.Now().Add(-1*24*time.Hour)) {
			klog.Info(v.GetName(), v.GetNamespace(), v.UnstructuredContent())
			client.Resource(pytorchRes).Namespace(v.GetNamespace()).Delete(context.TODO(), v.GetName(), metav1.DeleteOptions{})
		}
	}
}

func ListApiResource(client dynamic.Interface) {
	resource := client.Resource(schema.GroupVersionResource{})
	list, err := resource.List(nil, metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(list)
}
