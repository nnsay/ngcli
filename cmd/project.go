package cmd

import (
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "The project theme",
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
