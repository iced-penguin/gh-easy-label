package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "easy-label",
	Short: "Easy label management",
	Long:  "Manage your github labels easily",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(NewEditCmd())
	rootCmd.AddCommand(NewListCmd())
	rootCmd.AddCommand(NewApplyCmd())
}
