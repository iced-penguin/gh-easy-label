package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := newRootCmd()
	rootCmd.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		fmt.Println(err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "easy-label",
		Short: "Easy label management",
		Long:  "Manage your github labels easily",
	}
	rootCmd.AddCommand(newEditCmd())
	rootCmd.AddCommand(newListCmd())
	rootCmd.AddCommand(newApplyCmd())
	return rootCmd
}
