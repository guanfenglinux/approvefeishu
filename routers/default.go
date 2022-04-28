package routers

import (
	"approvefeishu/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.JenkinsController{})
}