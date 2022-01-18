package lib

import (
	"io"
	"net/http"
	"strconv"
	"time"
)

var fetch Fetch
var c = &http.Client{
	Timeout: 30 * time.Second,
}

func doRequest(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	config, err := ReadConfig()
	if err == nil && config.Auth.Token != "" {
		req.Header.Set("x-ng-application-type", strconv.Itoa(config.ApplicationType))
		req.Header.Set("x-ng-orgid", strconv.Itoa(config.Auth.OrgId))
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

// Fetch -- http require helper
type Fetch struct {
	Request    func(method string, url string, body io.Reader) ([]byte, error)
	SetTimeout func(timeout time.Duration)
}

func init() {
	fetch = Fetch{
		Request: doRequest,
		SetTimeout: func(_timeout time.Duration) {
			c = &http.Client{
				Timeout: _timeout,
			}
		},
	}
}

// GetFetch -- factory for get the http request helper
func GetFetch() Fetch {
	return fetch
}
