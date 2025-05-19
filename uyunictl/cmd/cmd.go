package cmd

import (
	"github.com/mmmmmmpc/uyuni-tools/shared/types"
	"github.com/mmmmmmpc/uyuni-tools/uyunictl/cmd/cp"
	"github.com/mmmmmmpc/uyuni-tools/uyunictl/cmd/exec"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command implementing the root command for kinder
func NewUyunictlCommand() *cobra.Command {
	globalFlags := &types.GlobalFlags{}
	rootCmd := &cobra.Command{
		Use:     "uyunictl",
		Short:   "Uyuni control tool",
		Long:    "Uyuni control tool used to help user managing Uyuni and SUSE Manager Servers mainly through its API",
		Version: "0.0.1",
	}

	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(exec.NewCommand(globalFlags))
	rootCmd.AddCommand(cp.NewCommand(globalFlags))

	return rootCmd
}
