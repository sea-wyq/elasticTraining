package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	gauge := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nvidia_gpu_allocated",
		Help: "nvidia gpu usage",
	})

	go func() {
		for {
			// Simulating memory usage update
			time.Sleep(time.Second)
			gauge.Set(float64(getGpuAllocatedUsage()))
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":7777", nil)
}

func getGpuAllocatedUsage() int {
	// config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{
		FieldSelector: "status.phase=Running",
	})
	if err != nil {
		log.Fatal(err)
	}

	totalUsage := 0
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			if _, ok := container.Resources.Requests["nvidia.com/nvidia-rtx-3090"]; ok {
				quantity := container.Resources.Requests["nvidia.com/nvidia-rtx-3090"]
				value := quantity.Value()
				totalUsage += int(value)
			}
		}
	}
	return totalUsage
}
