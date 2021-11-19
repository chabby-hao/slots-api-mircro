package user_api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type (
	ApiClient struct {
		conf       *UserApiConfig
		httpClient *http.Client
	}
)

var (
	ErrorUserNotFound = errors.New("User not found")
)

func NewUserApiClient(conf *UserApiConfig) *ApiClient {
	return &ApiClient{
		conf: conf,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (client *ApiClient) apiGet(path string, params map[string]interface{}) ([]byte, error) {
	queries := url.Values{}
	for k, v := range params {
		queries.Add(k, fmt.Sprintf("%v", v))
	}
	queries.Set("token", client.conf.Token)
	u := client.conf.BaseUrl + path + "?" + queries.Encode()
	logx.Infof("user api call: %s", u)
	res, err := client.httpClient.Get(u)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("user api get status code is %v", res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user api get read body error: %+v", err.Error()))
	}
	logx.Infof("res of %s: %s", u, string(body))
	return body, nil
}

func (client *ApiClient) SetUserPromoteTag(id1, id2, afid, tag string) error {
	rs, err := client.apiGet("/set-tenant-user", map[string]interface{}{
		"id1":    id1,
		"id2":    id2,
		"afid":   afid,
		"tenant": tag,
	})
	if err != nil {
		return err
	}
	var rsD ApiResult
	err = json.Unmarshal(rs, &rsD)
	if err != nil {
		return errors.New(fmt.Sprintf("parse api result failed: %+v", err.Error()))
	}
	if rsD.Code == 40100 {
		return ErrorUserNotFound
	}
	if rsD.Code != 0 {
		return errors.New(fmt.Sprintf("api rs code error. code: %d, msg: %s", rsD.Code, rsD.Message))
	}

	return nil
}
