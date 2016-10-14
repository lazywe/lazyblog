package admin

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	Time uint
	beego.Controller
}

//
// 初始化函数
//
func (this *BaseController) Prepare() {
	beego.SetViewsPath(beego.AppConfig.String("admin.tpl"))
	this.Time = uint(time.Now().Unix())
	// btnid := this.Input().Get("btnid")
	// intid, _ := strconv.Atoi(btnid)
	// if intid == 0 {
	// 	this.Abort("404")
	// }
	fmt.Println("这里是数据验证")
}

//
// Ajax返回信息
//
func (this *BaseController) AjaxReturn(status string, info string, data interface{}) {
	this.Data["json"] = map[string]interface{}{"status": status, "info": info, "data": data}
	this.ServeJSON()
}
