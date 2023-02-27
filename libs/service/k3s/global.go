package k3s_service

import (
	"flag"
	"testing"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var clientSet *kubernetes.Clientset

func init() {
	testing.Init()

	config, err := getConfig()
	if err != nil {
		panic(err)
	}

	clientS, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	clientSet = clientS
}

func getConfig() (*rest.Config, error) {
	// Pod内获取Config
	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}

	// Pod外获取Config
	kubeconfig := flag.String("kubeconfig", "/etc/rancher/k3s/k3s.yaml", "path to the kubeconfig file")
	flag.Parse()

	return clientcmd.BuildConfigFromFlags("", *kubeconfig)
}

func GetClientSet() *kubernetes.Clientset {
	return clientSet
}
