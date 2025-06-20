package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tutran-se/permiflow/internal/permiflow"
)

var (
	allowedVerbs      string
	excludedResources string
	outputPath        string
	roleName          string
	roleKind          string
	namespace         string
	scopeFilter       string
	explainOnly       bool
	profilePreset     string
)

var generateCmd = &cobra.Command{
	Use:   "generate-role",
	Short: "Generate a safe Kubernetes ClusterRole or Role",
	Long: `Generate a ClusterRole or Role that grants broad access while excluding dangerous or sensitive resources.
Useful for building secure default roles for bots, staging clusters, or external contractors.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if profilePreset != "" {
			if err := applyProfile(profilePreset); err != nil {
				return err
			}
		}

		verbs := strings.Split(allowedVerbs, ",")
		excludes := strings.Split(excludedResources, ",")

		if scopeFilter == "namespaced" {
			excludes = append(excludes,
				"nodes", "persistentvolumes", "clusterroles", "clusterrolebindings",
				"customresourcedefinitions", "mutatingwebhookconfigurations", "validatingwebhookconfigurations",
			)
		}

		opts := permiflow.GenerateRBACRoleOptions{
			AllowedVerbs:      clean(verbs),
			ExcludedResources: clean(excludes),
			OutputPath:        outputPath,
			RoleName:          roleName,
			RoleKind:          roleKind,
			Namespace:         namespace,
		}

		if strings.ToLower(roleKind) == "role" && namespace == "" {
			return fmt.Errorf("--namespace is required when --kind=Role")
		}

		if explainOnly {
			fmt.Println("\nSummary Preview:")
			fmt.Printf("- Verbs: %v\n", opts.AllowedVerbs)
			fmt.Printf("- Excluded Resources: %v\n", opts.ExcludedResources)
			fmt.Printf("- Role Kind: %s\n", opts.RoleKind)
			if opts.RoleKind == "Role" {
				fmt.Printf("- Namespace: %s\n", opts.Namespace)
			}
			fmt.Printf("- Role Name: %s\n", opts.RoleName)
			fmt.Printf("- Output Path: %s\n", opts.OutputPath)
			fmt.Println("(Use --dry-run to preview YAML, or omit to write to disk)")
			return nil
		}

		if dryRun {
			log.Println("Dry run enabled — printing YAML output only:")
			return permiflow.GenerateRBACRoleToStdout(opts)
		}

		return permiflow.GenerateRBACRole(opts)
	},
}

func clean(items []string) []string {
	out := []string{}
	for _, i := range items {
		i = strings.TrimSpace(i)
		if i != "" {
			out = append(out, i)
		}
	}
	return out
}

func applyProfile(profile string) error {
	switch profile {
	case "safe-cluster-admin":
		allowedVerbs = "get,list,watch,create,update"
		excludedResources = "secrets,pods/exec,nodes,rolebindings,clusterroles"
		roleName = "safe-cluster-admin"
		roleKind = "ClusterRole"
	case "read-only":
		allowedVerbs = "get,list,watch"
		excludedResources = "secrets,pods/exec,nodes,create,update,delete"
		roleName = "read-only-access"
		roleKind = "ClusterRole"
	default:
		return fmt.Errorf("unknown profile: %s", profile)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVar(&allowedVerbs, "allow-verbs", "get,list,watch", "Comma-separated verbs to allow")
	generateCmd.Flags().StringVar(&excludedResources, "exclude-resources", "secrets,pods/exec", "Comma-separated resources to exclude")
	generateCmd.Flags().StringVar(&scopeFilter, "scope", "", "Filter resource scope: namespaced or cluster")
	generateCmd.Flags().StringVar(&outputPath, "out", "clusterrole.yaml", "Output file path")
	generateCmd.Flags().StringVar(&roleName, "name", "almost-admin", "Name of the generated Role or ClusterRole")
	generateCmd.Flags().StringVar(&roleKind, "kind", "ClusterRole", "Kind to generate: ClusterRole or Role")
	generateCmd.Flags().StringVar(&namespace, "namespace", "", "Namespace for Role (required if kind=Role)")
	generateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print YAML to stdout instead of writing file")
	generateCmd.Flags().BoolVar(&explainOnly, "explain", false, "Show summary without generating role")
	generateCmd.Flags().StringVar(&profilePreset, "profile", "", "Use a predefined profile (safe-cluster-admin, read-only)")
}
