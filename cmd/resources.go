package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tutran-se/permiflow/internal/permiflow"
)

var (
	groupFilter    string
	versionFilter  string
	namespacedOnly bool
	outputJSON     bool
)

var resourcesCmd = &cobra.Command{
	Use:   "resources",
	Short: "List Kubernetes API resources and their group/version",
	Long:  "Shows available Kubernetes resource kinds grouped by API group/version. Supports filtering and JSON output.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := permiflow.GetKubeClient(kubeconfig)

		resources, err := permiflow.ListAPIResources(client.Discovery(), groupFilter, versionFilter, namespacedOnly)
		if err != nil {
			return err
		}

		if outputJSON {
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			return enc.Encode(resources)
		}

		// Human-friendly grouped output
		groupMap := make(map[string][]permiflow.ResourceInfo)
		for _, r := range resources {
			groupMap[r.GroupVersion] = append(groupMap[r.GroupVersion], r)
		}

		// Enhanced UX output
		fmt.Println("Kubernetes API Resources by GroupVersion:")
		fmt.Println("------------------------------------------------")

		groupVersions := make([]string, 0, len(groupMap))
		for gv := range groupMap {
			groupVersions = append(groupVersions, gv)
		}
		sort.Strings(groupVersions)

		for _, gv := range groupVersions {
			fmt.Printf("\n=== Group: %s ===\n", gv)
			resList := groupMap[gv]
			sort.Slice(resList, func(i, j int) bool {
				return resList[i].Name < resList[j].Name
			})

			for _, res := range resList {
				scope := "cluster-wide"
				if res.Namespaced {
					scope = "namespaced"
				}

				verbs := strings.Join(res.Verbs, ", ")
				fmt.Printf("  • %-30s  [scope: %s]  verbs: [%s]\n", res.Name, scope, verbs)
			}
		}

		fmt.Println("\nTip: Want to generate a secure ClusterRole?")
		fmt.Println("   permiflow generate-role --allow-verbs get,list --exclude-resources secrets,pods/exec")
		fmt.Println()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(resourcesCmd)

	resourcesCmd.Flags().StringVar(&groupFilter, "group", "", "Filter by API group (e.g. 'apps')")
	resourcesCmd.Flags().StringVar(&versionFilter, "version", "", "Filter by API version (e.g. 'v1')")
	resourcesCmd.Flags().BoolVar(&namespacedOnly, "namespaced-only", false, "Only show namespaced resources")
	resourcesCmd.Flags().BoolVar(&outputJSON, "json", false, "Output as JSON")
}
