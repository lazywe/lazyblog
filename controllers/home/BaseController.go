package home

import (
	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	beego.SetViewsPath(beego.AppConfig.String("home.tpl"))
	// btnid := this.Input().Get("btnid")
	// intid, _ := strconv.Atoi(btnid)
	// if intid == 0 {
	// 	this.Abort("404")
	// }
	fmt.Println("这里是数据验证")
}
