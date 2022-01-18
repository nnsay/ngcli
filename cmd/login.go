/*
Copyright © 2022 Jimmy Wang <jimmy.w@aliyun.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nnsay/ngcli/lib"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// LoginDTO -- login api request dto
type LoginDTO struct {
	ApplicationType int    `json:"applicationType"`
	UserName        string `json:"email"`
	Password        string `json:"password"`
}

// LoinResultDTO -- login api response dto
type LoinResultDTO struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	User    struct {
		ID    int `json:"id"`
		OrgID int `json:"orgId"`
	} `json:user`
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with username/password which are prividers by flags or configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		username := viper.GetString("username")
		password := viper.GetString("password")
		applicationType := viper.GetInt("applicationType")
		loginDTO := LoginDTO{applicationType, username, password}
		body, _ := json.Marshal(loginDTO)
		url := fmt.Sprintf("https://%s/%s", viper.GetString("endpoint"), lib.API_AUTH_LOGIN)

		result, err := lib.GetFetch().Request(http.MethodPost, url, bytes.NewBuffer(body))
		if err != nil {
			log.Panic(err)
		}
		loginResult := LoinResultDTO{}
		err = json.Unmarshal(result, &loginResult)
		if err != nil {
			log.Panic(err)
		}

		config, err := lib.ReadConfig()
		if err != nil {
			log.Panic(err)
		}
		config.Auth = lib.AuthConfig{
			Token:   loginResult.Token,
			TokenAt: time.Now().Format("2006-01-02 13:04:06"),
			OrgId:   loginResult.User.OrgID,
		}
		lib.SaveConfig(config)

		fmt.Println("Login Success")
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
}
