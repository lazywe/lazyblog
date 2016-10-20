package home

import (
	"lazyblog/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

var Menus []models.Menu

func (this *BaseController) Prepare() {
	beego.SetViewsPath(beego.AppConfig.String("home.tpl"))
	// btnid := this.Input().Get("btnid")
	// intid, _ := strconv.Atoi(btnid)
	// if intid == 0 {
	// 	this.Abort("404")
	// }
	// fmt.Println("这里是数据验证")

	//导航
	var menuMode = new(models.Menu)
	_, Menus = menuMode.GetHomeMenuList()
	this.Data["Menus"] = Menus
}
