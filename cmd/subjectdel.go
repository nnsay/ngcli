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

// subjectdelCmd represents the subjectdel command
var subjectdelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a subject",
	Run: func(cmd *cobra.Command, args []string) {
		subjectId := viper.GetInt("subjectId")
		projectId := viper.GetInt("projectId")
		url := fmt.Sprintf("https://%s/%s/%d/%d", viper.GetString("endpoint"), lib.API_SUBJECT, subjectId, projectId)
		byteBody, err := lib.GetFetch().Request(http.MethodDelete, url, nil)
		if err != nil {
			log.Panic(err)
		}
		prettyJSON, _ := lib.PrettyJSON(byteBody)
		fmt.Println(prettyJSON)
	},
}

func init() {
	subjectCmd.AddCommand(subjectdelCmd)

	subjectdelCmd.Flags().Int("subjectId", 0, "required, a subject id, eg: 3959")
	subjectdelCmd.MarkFlagRequired("subjectId")
	viper.BindPFlag("subjectId", subjectdelCmd.Flags().Lookup("subjectId"))

	subjectdelCmd.Flags().Int("projectId", 0, "required, a project id, eg: 207")
	subjectdelCmd.MarkFlagRequired("projectId")
	viper.BindPFlag("projectId", subjectdelCmd.Flags().Lookup("projectId"))
}
