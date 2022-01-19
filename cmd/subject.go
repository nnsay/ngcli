/*
Copyright © 2022 Jimmy Wang <jimmy.w@aliyun.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// subjectCmd represents the subject command
var subjectCmd = &cobra.Command{
	Use:   "subject",
	Short: "The subject theme",
}

func init() {
	rootCmd.AddCommand(subjectCmd)
}
