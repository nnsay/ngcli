/*
Copyright © 2022 Jimmy Wang <jimmy.w@aliyun.com>

*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"nnsay/ngcli/lib"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// subjectlistCmd represents the subjectlist command
var subjectlistCmd = &cobra.Command{
	Use:   "list",
	Short: "List subject of current user org",
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("https://%s/%s", viper.GetString("endpoint"), lib.API_SUBJECT)
		byteBody, err := lib.GetFetch().Request(http.MethodGet, url, nil)
		if err != nil {
			log.Panic(err)
		}
		prettyJSON, _ := lib.PrettyJSON(byteBody)
		fmt.Println(prettyJSON)
	},
}

func init() {
	subjectCmd.AddCommand(subjectlistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subjectlistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subjectlistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
