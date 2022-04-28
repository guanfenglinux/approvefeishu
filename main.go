package main

import (
	_ "approvefeishu/routers"
	_ "approvefeishu/structs"
	"github.com/astaxie/beego"
)

func main() {
	addr := beego.AppConfig.String("server::addr")
	beego.Run(addr)
}
