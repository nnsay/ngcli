/*
Copyright © 2022 Jimmy Wang <jimmy.w@aliyun.com>

*/
package cmd

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"nnsay/ngcli/lib"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// subjectcreateCmd represents the subjectadd command
var subjectcreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new subject",
	Run: func(cmd *cobra.Command, args []string) {
		subjectCustId, _ := cmd.Flags().GetString("subjectCustId")
		projectId, _ := cmd.Flags().GetInt("projectId")

		url := fmt.Sprintf("https://%s/%s", viper.GetString("endpoint"), lib.API_SUBJECT)
		data := []byte(fmt.Sprintf(`{"subjectCustId":"%s","projectId":%d}`, subjectCustId, projectId))
		byteBody, err := lib.GetFetch().Request(http.MethodPost, url, bytes.NewBuffer(data))
		if err != nil {
			log.Fatal(err)
		}
		prettyJSON, _ := lib.PrettyJSON(byteBody)
		fmt.Println(prettyJSON)
	},
}

func init() {
	subjectCmd.AddCommand(subjectcreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	subjectcreateCmd.Flags().StringP("subjectCustId", "s", "", "required, subject custome id, eg: sub001")
	subjectcreateCmd.MarkFlagRequired("subjectCustId")

	subjectcreateCmd.Flags().Int("projectId", 0, "required, project id, it can be get from project list command")
	subjectcreateCmd.MarkFlagRequired("projectId")
}
