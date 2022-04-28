package services

import (
	"approvefeishu/structs"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"net/http"
)

type Build struct {
	UserId       string
	OpenId       string
	InstanceCode string
}

// GetApproveInfo 获取审核实例详情
func (b *Build) GetApproveInfo() (string, string, error) {
	var name, tag string
	token, err := GetToken()
	url := "https://www.feishu.cn/approval/openapi/v2/instance/get"
	request := structs.InstanceInfo{
		InstanceCode: b.InstanceCode,
		Locale:       "zh-CN",
		UserID:       b.UserId,
		OpenID:       b.OpenId,
	}

	body, _ := json.Marshal(request)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logs.Error(err)
			return
		}
	}(resp.Body)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	var rep structs.Instance
	var form structs.Form
	err = json.Unmarshal(data, &rep)
	if err != nil {
		return "", "", err
	}
	dataForm := rep.Data.Form
	fmt.Println(dataForm)
	err = json.Unmarshal([]byte(dataForm), &form)
	if err != nil {
		return "", "", err
	}

	for _, v := range form {
		if v.Name == "上线项目" {
			name = v.Value
		}
		if v.Name == "版本tag" {
			tag = v.Value
		}
	}
	return name, tag, nil
}
