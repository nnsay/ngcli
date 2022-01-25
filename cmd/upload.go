/*
Copyright © 2022 Jimmy Wang <jimmy.w@aliyun.com>

*/
package cmd

import (
	"bytes"
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
		partSize := viper.GetInt64("partSize")
		parallelNumber := viper.GetInt("parallelNumber")
		projectId := viper.GetInt("projectId")

		stsToken := uploadSTSToken(subjectId, fileName)

		upload(file, stsToken, partSize, parallelNumber)
		uploadId := createUpload(projectId, subjectId, fileName, stsToken.OSSUri)

		submitUpload(uploadId, subjectId)

		jobId := createJob(projectId, subjectId, uploadId)

		submitJob(jobId, subjectId)

		fmt.Println(generateWebUrl(jobId, subjectId))
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

	uploadCmd.Flags().Int64P("partSize", "b", 0, "optional, the upload part size, default: 5242880")
	viper.BindPFlag("partSize", uploadCmd.Flags().Lookup("partSize"))
	viper.SetDefault("partSize", 5242880)

	uploadCmd.Flags().IntP("parallelNumber", "n", 0, "optional, paralle upload process count, default: 5")
	viper.BindPFlag("parallelNumber", uploadCmd.Flags().Lookup("parallelNumber"))
	viper.SetDefault("parallelNumber", 5)

	uploadCmd.Flags().Int("projectId", 0, "required, project id, it can be get from project list command")
	uploadCmd.MarkFlagRequired("projectId")
	viper.BindPFlag("projectId", uploadCmd.Flags().Lookup("projectId"))
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

func upload(file string, stsToken types.UploadSTSDTO, partSize int64, parallelNumber int) {
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

	err = bucket.UploadFile(stsToken.OSSKey, file, partSize, oss.Routines(parallelNumber))
	if err != nil {
		log.Panic(err)
	}
}

func createUpload(projectId int, subjectId int, fileName string, s3uri string) int {
	jsonStr := fmt.Sprintf(`{"ln":"%s","mimeType":"application/zip","s3uri":"%s","projectId":%d}`, fileName, s3uri, projectId)
	data := []byte(jsonStr)
	url := fmt.Sprintf("https://%s/%s/%d/uploads", viper.GetString("endpoint"), lib.API_SUBJECT, subjectId)
	byteBody, err := lib.GetFetch().Request(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		log.Panic(err)
	}
	createdUpload := types.CreatedUploadDTO{}
	json.Unmarshal(byteBody, &createdUpload)

	return createdUpload.Upload.ID
}

func submitUpload(uploadId int, subjectId int) {
	jsonStr := fmt.Sprintf(`{"id":%d,"uploadStatus":%d}`, uploadId, lib.UploadStatusReady)
	data := []byte(jsonStr)
	url := fmt.Sprintf("https://%s/%s/%d/uploads/%d", viper.GetString("endpoint"), lib.API_SUBJECT, subjectId, uploadId)
	_, err := lib.GetFetch().Request(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		log.Panic(err)
	}
}

func createJob(projectId int, subjectId int, uploadId int) int {
	// HINT: hardcode the jobType for StepGroupProcess
	jsonStr := fmt.Sprintf(`{"jobType":4,"force":true,"uploadId":%d,"projectId":%d,"labels":[]}`, uploadId, projectId)
	data := []byte(jsonStr)
	url := fmt.Sprintf("https://%s/%s/%d/jobs", viper.GetString("endpoint"), lib.API_SUBJECT, subjectId)
	byteBody, err := lib.GetFetch().Request(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		log.Panic(err)
	}
	createdJob := types.CreatedJobDTO{}
	json.Unmarshal(byteBody, &createdJob)

	return createdJob.ID
}

func submitJob(jobId int, subjectId int) {
	data := []byte("{}")
	url := fmt.Sprintf("https://%s/%s/%d/jobs/%d/submit", viper.GetString("endpoint"), lib.API_SUBJECT, subjectId, jobId)
	_, err := lib.GetFetch().Request(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		log.Panic(err)
	}
}

func generateWebUrl(jobId, subjectId int) string {
	webDomain := strings.Replace(viper.GetString("endpoint"), "-api.", ".", 1)
	return fmt.Sprintf(`Upload Success!
	to view in browser: https://%s/#/dataservice/subject/%d/info/job/%d/files?detail=false
	`, webDomain, subjectId, jobId)
}
