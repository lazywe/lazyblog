package routers

import (
	"lazyblog/controllers/admin" //后台
	"lazyblog/controllers/home"  //前台

	"github.com/astaxie/beego"
)

func init() {
	//后台
	beego.Router("/main", &admin.AdminController{})
	beego.Router("/main/panal", &admin.AdminController{}, "get:Panal")
	beego.Router("/main/login", &admin.LoginController{}, "get:Login")
	beego.Router("/main/login", &admin.LoginController{}, "post:LoginDo")
	beego.Router("/main/node/list.html", &admin.NodeController{}, "get:Node")

	//功能类型
	beego.Router("/main/option/option.html", &admin.OptionController{}, "get:Option")
	beego.Router("/main/option/addoption.html", &admin.OptionController{}, "get:AddOption")
	beego.Router("/main/option/addoption.html", &admin.OptionController{}, "post:AddOptionDo")
	beego.Router("/main/option/editoption.html", &admin.OptionController{}, "get:EditOption")
	beego.Router("/main/option/editoption.html", &admin.OptionController{}, "post:EditOptionDo")
	beego.Router("/main/option/deloption.html", &admin.OptionController{}, "get:DelOption")

	//前台
	beego.Router("/", &home.IndexController{})
}
