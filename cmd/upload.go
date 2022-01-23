/*
Copyright © 2022 Jimmy Wang <jimmy.w@aliyun.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nnsay/ngcli/lib"
	"nnsay/ngcli/types"
	"path/filepath"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload file(zip)",

	Run: func(cmd *cobra.Command, args []string) {
		file := viper.GetString("file")
		subjectId := viper.GetInt("subjectId")
		fileName := filepath.Base(file)

		stsToken := uploadSTSToken(subjectId, fileName)
		upload(file, stsToken)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringP("file", "f", "", "required, the upload zip file path, eg: ./philips-reconall.zip")
	uploadCmd.MarkFlagRequired("file")
	viper.BindPFlag("file", uploadCmd.Flags().Lookup("file"))

	uploadCmd.Flags().IntP("subjectId", "s", 0, "required, the subject id, eg: 1268")
	uploadCmd.MarkFlagRequired("subjectId")
	viper.BindPFlag("subjectId", uploadCmd.Flags().Lookup("subjectId"))
}

func uploadSTSToken(subjectId int, fileName string) types.UploadSTSDTO {
	url := fmt.Sprintf("https://%s/%s/%d/uploads/ossuploadtoken?ln=%s&mimeType=application/octet-stream", viper.GetString("endpoint"), lib.API_SUBJECT, subjectId, fileName)
	byteBody, err := lib.GetFetch().Request(http.MethodGet, url, nil)
	if err != nil {
		log.Panic(err)
	}

	accessKey := types.UploadSTSDTO{}
	json.Unmarshal(byteBody, &accessKey)
	return accessKey
}

func upload(file string, stsToken types.UploadSTSDTO) {
	client, err := oss.New(
		fmt.Sprintf("%s.aliyuncs.com", stsToken.OSSRegion),
		stsToken.AccessKeyId,
		stsToken.AccessKeySecret,
		oss.SecurityToken(stsToken.SecurityToken),
	)
	if err != nil {
		log.Panic(err)
	}
	bucktName := strings.Split(strings.Split(stsToken.OSSUri, "//")[1], "/")[0]
	bucket, err := client.Bucket(bucktName)
	if err != nil {
		log.Panic(err)
	}
	err = bucket.PutObjectFromFile(stsToken.OSSUri, file)
	if err != nil {
		log.Panic(err)
	}
}
