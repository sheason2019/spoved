package k3s_service

import (
	"k8s.io/client-go/kubernetes"
)

var clientSet *kubernetes.Clientset

func init() {
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
