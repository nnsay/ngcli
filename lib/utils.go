package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

func PrettyJSON(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MakeSureConfigFile(file string) {
	isExist, _ := PathExists(file)
	if !isExist {
		ioutil.WriteFile(file, []byte("# created by ngcli"), 0755)
	}
}
