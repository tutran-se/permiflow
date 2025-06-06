package main

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeClient() *kubernetes.Clientset {
  home, err := os.UserHomeDir()
  if err != nil {
    log.Fatalf("Cannot find user home directory: %v", err)
  }
  kubeconfig := filepath.Join(home, ".kube", "config")

  config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
  if err != nil {
    log.Fatalf("Error loading kubeconfig from %s: %v", kubeconfig, err)
  }

  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    log.Fatalf("Error creating clientset: %v", err)
  }

  return clientset
}
