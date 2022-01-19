package lib

import (
	"bytes"
	"encoding/json"
)

func PrettyJSON(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
