package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"nnsay/ngcli/types"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type Fetch struct {
	c *http.Client
}

var fetch Fetch

func (f *Fetch) Request(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	config, err := ReadConfig()
	if err == nil && config.Auth.Token != "" {
		req.Header.Set("x-ng-application-type", strconv.Itoa(config.ApplicationType))
		req.Header.Set("x-ng-orgid", strconv.Itoa(config.Auth.OrgId))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Auth.Token))
	}

	res, err := f.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= http.StatusBadRequest {
		erroCode := types.ErrorCode{}
		err = json.Unmarshal(bodyBytes, &erroCode)
		if err != nil {
			return nil, err
		}
		// if the errcode is auth failure retry,
		// the normal auth failure reason is JWT token expiration
		if erroCode.Code == HTTP_ERROR_AUTHFAILED {
			err = f.Login()
			if err != nil {
				return nil, err
			}
			return f.Request(method, url, body)
		}
		return nil, errors.New(string(bodyBytes))
	}
	return bodyBytes, nil
}

func (f *Fetch) SetTimeout(_timeout time.Duration) {
	f.c = &http.Client{
		Timeout: _timeout,
	}
}

func (f *Fetch) Login() error {
	username := viper.GetString("username")
	password := viper.GetString("password")
	applicationType := viper.GetInt("applicationType")
	loginDTO := types.LoginDTO{ApplicationType: applicationType, UserName: username, Password: password}
	body, _ := json.Marshal(loginDTO)
	url := fmt.Sprintf("https://%s/%s", viper.GetString("endpoint"), API_AUTH_LOGIN)

	result, err := fetch.Request(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	loginResult := types.LoinResultDTO{}
	err = json.Unmarshal(result, &loginResult)
	if err != nil {
		return err
	}

	config, err := ReadConfig()
	if err != nil {
		return err
	}
	config.Auth = types.AuthConfig{
		Token:   loginResult.Token,
		TokenAt: time.Now().Format("2006-01-02 15:04:05"),
		OrgId:   loginResult.User.OrgID,
	}
	err = SaveConfig(config)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	fetch = Fetch{
		c: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetFetch -- factory for get the http request helper
func GetFetch() *Fetch {
	return &fetch
}
