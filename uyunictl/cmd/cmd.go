package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uyuni-project/mmmmmmpc/shared/types"
	"github.com/uyuni-project/mmmmmmpc/uyunictl/cmd/cp"
	"github.com/uyuni-project/mmmmmmpc/uyunictl/cmd/exec"
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
