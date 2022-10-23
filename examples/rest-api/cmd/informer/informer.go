package main

import (
	"flag"
	"fmt"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/informers"
	jobformers "k8s.io/client-go/informers/batch/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	klog "k8s.io/klog/v2"
	//"k8s.io/client-go/pkg/api/v1"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/component-base/logs"
)

// PodLoggingController logs the name and namespace of pods that are added,
// deleted, or updated
type PodLoggingController struct {
	informerFactory informers.SharedInformerFactory
	podInformer     coreinformers.PodInformer
	jobInformer     jobformers.JobInformer
}

// Run starts shared informers and waits for the shared informer cache to
// synchronize.
func (c *PodLoggingController) Run(stopCh chan struct{}) error {
	// Starts all the shared informers that have been created by the factory so
	// far.
	c.informerFactory.Start(stopCh)
	// wait for the initial synchronization of the local cache.
	if !cache.WaitForCacheSync(stopCh, c.podInformer.Informer().HasSynced) {
		return fmt.Errorf("Failed to sync")
	}
	return nil
}

func (c *PodLoggingController) podAdd(obj interface{}) {
	pod := obj.(*v1.Pod)
	klog.Infof("POD CREATED: %s/%s", pod.Namespace, pod.Name)
}

func (c *PodLoggingController) podUpdate(old, new interface{}) {
	oldPod := old.(*v1.Pod)
	newPod := new.(*v1.Pod)
	klog.Infof(
		"POD UPDATED. %s/%s %s",
		oldPod.Namespace, oldPod.Name, newPod.Status.Phase,
	)
}

func (c *PodLoggingController) podDelete(obj interface{}) {
	job := obj.(*v1.Pod)
	klog.Infof("POD DELETED: %s/%s", job.Namespace, job.Name)
}

// JOB informer
func (c *PodLoggingController) jobAdd(obj interface{}) {
	job := obj.(*batchv1.Job)
	klog.Infof("JOB CREATED: %s/%s", job.Namespace, job.Name)
}

func (c *PodLoggingController) jobUpdate(old, new interface{}) {
	oldJob := old.(*batchv1.Job)
	newJob := new.(*batchv1.Job)
	klog.Infof(
		"JOB UPDATED. %s/%s %s",
		oldJob.Namespace, oldJob.Name, newJob.TypeMeta.String(),
	)
}

func (c *PodLoggingController) jobDelete(obj interface{}) {
	pod := obj.(*batchv1.Job)
	klog.Infof("JOB DELETED: %s/%s", pod.Namespace, pod.Name)
}

// NewPodLoggingController creates a PodLoggingController
func NewPodLoggingController(informerFactory informers.SharedInformerFactory) *PodLoggingController {
	podInformer := informerFactory.Core().V1().Pods()
	jobInformer := informerFactory.Batch().V1().Jobs()

	c := &PodLoggingController{
		informerFactory: informerFactory,
		podInformer:     podInformer,
		jobInformer:     jobInformer,
	}
	podInformer.Informer().AddEventHandler(
		// Your custom resource event handlers.
		cache.ResourceEventHandlerFuncs{
			// Called on creation
			AddFunc: c.podAdd,
			// Called on resource update and every resyncPeriod on existing resources.
			UpdateFunc: c.podUpdate,
			// Called on resource deletion.
			DeleteFunc: c.podDelete,
		},
	)

	jobInformer.Informer().AddEventHandler( // Your custom resource event handlers.
		cache.ResourceEventHandlerFuncs{
			// Called on creation
			AddFunc: c.jobAdd,
			// Called on resource update and every resyncPeriod on existing resources.
			UpdateFunc: c.jobUpdate,
			// Called on resource deletion.
			DeleteFunc: c.jobDelete,
		},
	)
	return c
}

var kubeconfig string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "/Users/wangbaoyi1/.kube/config", "absolute path to the kubeconfig file")
}

func main() {
	flag.Parse()
	logs.InitLogs()
	defer logs.FlushLogs()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	factory := informers.NewSharedInformerFactory(clientSet, time.Hour*24)
	controller := NewPodLoggingController(factory)
	stop := make(chan struct{})
	defer close(stop)
	err = controller.Run(stop)
	if err != nil {
		klog.Fatal(err)
	}
	select {}
}
