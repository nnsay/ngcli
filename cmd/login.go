/*
Copyright © 2022 Jimmy Wang <jimmy.w@aliyun.com>

*/
package cmd

import (
	"fmt"
	"log"
	"nnsay/ngcli/lib"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with username/password which are prividers by flags or configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		err := lib.GetFetch().Login()
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("Login Success")
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
}
