package services

import (
	"approvefeishu/structs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// GetToken 获取飞书token
func GetToken() (token string, err error) {
	rep := structs.ResponseToken{}
	r := structs.RequestToken{
		AppId:     "xxxxxxxx",
		AppSecret: "xxxxxxxxx",
	}
	body, _ := json.Marshal(r)
	client := &http.Client{}
	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &rep)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Bearer %s", rep.TenantAccessToken), nil
}
