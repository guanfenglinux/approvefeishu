package services

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/bndr/gojenkins"
)

func BuildJob(name, tag string) error {
	user := beego.AppConfig.String("jenkins::user")
	passwd := beego.AppConfig.String("jenkins::passwd")
	host := beego.AppConfig.String("jenkins::host")
	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, host, user, passwd)
	_, err := jenkins.Init(ctx)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["BRANCH_TAG"] = tag
	_, err = jenkins.BuildJob(ctx, name, params)
	if err != nil {
		return err
	}
	return nil
}
