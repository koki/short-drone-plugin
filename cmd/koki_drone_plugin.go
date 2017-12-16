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
		Long: `koki-short converts the koki manifests into Kubernetes syntax.

Full documentation available at https://docs.koki.io/short
`,
		RunE: func(c *cobra.Command, args []string) error {
			return executor.Execute(files.String(), outputPrefix.String(), inPlace.String())
		},
		Example: `

`,
	}
)

//flags for the plugin
var (
	inPlace      types.FlagOrEnv
	outputPrefix types.FlagOrEnv
	files        types.FlagOrEnv
)

const droneEnvPrefix = "PLUGIN_"

func init() {
	inPlace.AddToCobraCommand(KokiCmd, "in-place", "i", "translate the files in place", droneEnvPrefix, true)
	outputPrefix.AddToCobraCommand(KokiCmd, "output-prefix", "p", "prefix for the translated files", droneEnvPrefix, false)
	files.AddToCobraCommand(KokiCmd, "files", "f", "path to koki short files that need to be transformed", droneEnvPrefix, false)
}
