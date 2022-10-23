package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/examples/rest-api/cluster"
)

func main() {
	listNode()
	listPods("public-resource")
	listJobs("public-resource")
	//unMarshalPodJson("/Users/wby/code/cncf/k8s/client-go-wby/examples/rest-api/api/test.json")
}

func listNode() {
	client := cluster.NewClusterClient()
	list, err := client.CoreV1().Nodes().List(context.Background(), v1.ListOptions{
		LabelSelector: "mlops/resource-group=public-resource",
	})
	if err != nil {
		log.Fatalln("get node list error: ", err.Error())
	}
	for _, v := range list.Items {
		fmt.Println(v.GetName())
		fmt.Println(v.Status.Phase)
	}
}

func listPods(ns string) {
	client := cluster.NewClusterClient()
	list, err := client.CoreV1().Pods(ns).List(context.Background(), v1.ListOptions{})

	if err != nil {
		log.Fatalln("get pod list error: ", list)
	}

	for _, v := range list.Items {
		if v.GetName() == "svc-save-all-to-last-model-235-7223-1665739120356" {
			b, _ := json.Marshal(v)
			fmt.Println(b)
		}

		fmt.Println(v.GetName())
		fmt.Println(v.Status.Phase)
	}
}

func listJobs(ns string) {
	client := cluster.NewClusterClient()
	list, err := client.BatchV1().Jobs(ns).List(context.Background(), v1.ListOptions{})

	if err != nil {
		log.Fatalln("get jobs list error: ", list)
	}

	for _, v := range list.Items {
		if v.GetName() == "svc-save-all-to-last-model-235-7223-1665739120356" {
			b, _ := json.Marshal(v)
			fmt.Println(b)
		}

		fmt.Println(v.GetName(), " ", v.Annotations)
	}
}

func unMarshalPodJson(fileName string) {
	pod := v12.Pod{}
	b, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("read pod json file error: ", err)
	}

	err = json.Unmarshal(b, pod)
	if err != nil {
		log.Fatalln("unmarshal json error: ", err)
	}
}
