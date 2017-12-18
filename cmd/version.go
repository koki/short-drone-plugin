package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	GITCOMMIT = "HEAD"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version of short-drone-plugin",
		Run: func(*cobra.Command, []string) {
			fmt.Printf("koki/short-drone-plugin: %s\n", GITCOMMIT)
		},
	}
)
