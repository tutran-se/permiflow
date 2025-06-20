package permiflow

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
)

// ResourceInfo represents a Kubernetes resource kind and its metadata
type ResourceInfo struct {
	Group        string   `json:"group"`
	Version      string   `json:"version"`
	Name         string   `json:"name"`
	Namespaced   bool     `json:"namespaced"`
	Verbs        []string `json:"verbs"`
	GroupVersion string   `json:"groupVersion"`
}

// ListAPIResources returns a flat list of resources from the kube-apiserver
func ListAPIResources(client discovery.DiscoveryInterface, groupFilter, versionFilter string, namespacedOnly bool) ([]ResourceInfo, error) {
	disco := memory.NewMemCacheClient(client)

	apiResourceLists, err := disco.ServerPreferredResources()
	if err != nil {
		return nil, fmt.Errorf("error fetching API resource list: %w", err)
	}

	var out []ResourceInfo
	for _, list := range apiResourceLists {
		groupVersion := list.GroupVersion

		parsed, err := schema.ParseGroupVersion(groupVersion)
		if err != nil {
			continue
		}
		group := parsed.Group
		version := parsed.Version

		if groupFilter != "" && group != groupFilter {
			continue
		}
		if versionFilter != "" && version != versionFilter {
			continue
		}

		for _, res := range list.APIResources {
			if namespacedOnly && !res.Namespaced {
				continue
			}
			out = append(out, ResourceInfo{
				Group:        group,
				Version:      version,
				Name:         res.Name,
				Namespaced:   res.Namespaced,
				Verbs:        res.Verbs,
				GroupVersion: groupVersion,
			})
		}
	}

	return out, nil
}
