package controllers

import (
	"approvefeishu/services"
	"approvefeishu/structs"
	"approvefeishu/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type JenkinsController struct {
	beego.Controller
}

func (c *JenkinsController) Api() {
	var req structs.ApproveTrigger
	data := c.Ctx.Input.RequestBody
	fmt.Println(string(data))
	err := json.Unmarshal(data, &req)
	if err != nil {
		logs.Error(err)
		c.ServeJSON()
		return
	}
	// 判断uuid是否存在防止重复推送，并且状态是同意,并且审核的id匹配
	if utils.Add(req.UUID) && req.Event.Status == "APPROVED" && req.Event.ApprovalCode == "xxxxxxxxxxxxxx" {
		b := services.Build{
			InstanceCode: req.Event.InstanceCode,
			UserId:       req.Event.UserID,
			OpenId:       req.Event.OpenID,
		}
		// 获取项目name和tag
		name, tag, err := b.GetApproveInfo()
		if err != nil {
			logs.Error(err)
			c.ServeJSON()
			return
		}
		// 执行调用jenkins构建的函数
		err = services.BuildJob(name, tag)
		if err != nil {
			logs.Error(err)
			c.ServeJSON()
			return
		}
	}

	c.ServeJSON()
	return
}
