package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"nnsay/ngcli/lib"
)

// projectlistCmd represents the projectlist command
var projectlistCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects of current user org",
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("https://%s/%s", viper.GetString("endpoint"), lib.API_PROJECT_LIST)
		byteBody, err := lib.GetFetch().Request(http.MethodGet, url, nil)
		if err != nil {
			log.Panic(err)
		}
		prettyJSON, _ := lib.PrettyJSON(byteBody)
		fmt.Println(prettyJSON)
	},
}

func init() {
	projectCmd.AddCommand(projectlistCmd)
}
