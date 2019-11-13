package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	flag.Parse()
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := flag.String("kubeconfig", "/Users/huangweixiao/.kube/config", "kubeconfig file")
		if envvar := os.Getenv("KUBECONFIG"); len(envvar) > 0 {
			*kubeconfig = envvar
		}
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			log.WithError(err).Fatal("The kubeconfig cannot be loaded")
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.WithError(err).Fatal("kubernetes.NewForConfig")
	}

	pod, err := clientset.CoreV1().Pods("kube-system").Get("etcd-docker-desktop", metav1.GetOptions{})
	if err != nil {
		log.WithError(err).Fatal("Get pods")
	}
	log.Infoln("pod: ", pod)
}
