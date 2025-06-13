package permiflow

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeClient(kubeconfigPath string) *kubernetes.Clientset {
	// If no kubeconfig path is provided, use the default path
	if kubeconfigPath == "" {
		// Check KUBECONFIG environment variable first
		if kubeconfigEnv := os.Getenv("KUBECONFIG"); kubeconfigEnv != "" {
			kubeconfigPath = kubeconfigEnv
		} else {
			// Use default path: ~/.kube/config
			homeDir, err := os.UserHomeDir()
			if err != nil {
				log.Printf("Warning: Could not get user home directory: %v", err)
			} else {
				kubeconfigPath = filepath.Join(homeDir, ".kube", "config")
			}
		}
	}

	// Build config from the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Printf("Error loading kubeconfig from %s: %v", kubeconfigPath, err)
		return nil
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Error creating clientset: %v", err)
		return nil
	}

	log.Printf("Successfully connected to Kubernetes cluster using kubeconfig: %s", kubeconfigPath)
	return clientset
}

func GetCurrentContext(kubeconfig string) string {
	cfg, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		return "unknown"
	}
	return cfg.CurrentContext
}
