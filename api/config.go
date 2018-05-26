package api

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"flag"
	"path/filepath"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

var (
	client *kubernetes.Clientset
)

func init() {
	config, err := rest.InClusterConfig()
	if err != nil {
		client = loadClient()
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	client = clientset
}

func loadClient() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}
