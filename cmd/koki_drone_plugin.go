package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kubeciio/koki/executor"
	"github.com/kubeciio/koki/types"
)

var (
	KokiCmd = &cobra.Command{
		Use:   "koki-short",
		Short: "Convert Koki Kubernetes manifests to validated Kubernetes manifests",
		Long: `  _  _____  _  _____ 
 | |/ / _ \| |/ /_ _|
 | ' < (_) | ' < | | 
 |_|\_\___/|_|\_\___|

koki-short converts koki manifests into Kubernetes syntax.

Full documentation available at https://docs.koki.io/short
`,
		RunE: func(c *cobra.Command, args []string) error {
			return executor.Execute(files.String(), outputPrefix.String(), inPlace.String(), shortPath.String(), overwrite.String())
		},
		SilenceUsage: true,
		Example: `
# Provide input files using flag
koki -f test1.yaml -f test2.yaml

# Provide input files using environment variable
PLUGIN_FILES=test1.yaml,test2.yaml koki 

# Provide output file prefix using flag
koki -f test1.yaml -p k8s_

# Provide output file prefix using environment variable
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml koki

# Translate file in-place using flag
koki -f test1.yaml -i

# Translate file in-place using environment variable
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml PLUGIN_IN_PLACE=true koki

# Preserves the directory in which the original file was found
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=/path/to/dir/test1.yaml,test2.yaml PLUGIN_IN_PLACE=true koki
 >  output file will be created in /path/to/dir/ and in current directory

# Set overwrite=true if output file already exists (using env vars)
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml PLUGIN_IN_PLACE=true PLUGIN_OVERWRITE=true koki

# Set overwrite=true if output file already exists (using flag)
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml PLUGIN_IN_PLACE=true koki -w

`,
	}
)

//flags for the plugin
var (
	inPlace      types.FlagOrEnv
	outputPrefix types.FlagOrEnv
	files        types.FlagOrEnv
	shortPath    types.FlagOrEnv
	overwrite    types.FlagOrEnv
)

const droneEnvPrefix = "PLUGIN_"

func init() {
	inPlace.AddToCobraCommand(KokiCmd, "in-place", "i", "false", "translate the files in place", droneEnvPrefix, true)
	outputPrefix.AddToCobraCommand(KokiCmd, "output-prefix", "p", "kube_", "prefix for the translated files", droneEnvPrefix, false)
	files.AddToCobraCommand(KokiCmd, "files", "f", "", "path to koki short files that need to be transformed", droneEnvPrefix, false)
	shortPath.AddToCobraCommand(KokiCmd, "short-path", "s", "", "path to short binary", droneEnvPrefix, false)
	overwrite.AddToCobraCommand(KokiCmd, "overwrite", "w", "false", "overwrite existing files (fails by default)", droneEnvPrefix, true)
}
