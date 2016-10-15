package routers

import (
	"lazyblog/controllers/admin" //后台
	"lazyblog/controllers/home"  //前台

	"github.com/astaxie/beego"
)

func init() {
	//后台
	beego.Router("/main.html", &admin.AdminController{}, "get:Main")
	beego.Router("/main/panal.html", &admin.AdminController{}, "get:Panal")
	beego.Router("/main/login.html", &admin.LoginController{}, "get:Login")
	beego.Router("/main/login.html", &admin.LoginController{}, "post:LoginDo")
	beego.Router("/main/node/list.html", &admin.NodeController{}, "get:Node")

	//功能类型
	beego.Router("/main/option/option.html", &admin.OptionController{}, "get:Option")
	beego.Router("/main/option/addoption.html", &admin.OptionController{}, "get:AddOption")
	beego.Router("/main/option/addoption.html", &admin.OptionController{}, "post:AddOptionDo")
	beego.Router("/main/option/editoption.html", &admin.OptionController{}, "get:EditOption")
	beego.Router("/main/option/editoption.html", &admin.OptionController{}, "post:EditOptionDo")
	beego.Router("/main/option/deloption.html", &admin.OptionController{}, "post:DelOption")
	beego.Router("/main/option/sortoption.html", &admin.OptionController{}, "post:SortOption")

	//功能类型
	beego.Router("/main/blog/blog.html", &admin.BlogController{}, "get:Blog")
	beego.Router("/main/blog/addblog.html", &admin.BlogController{}, "get:AddBlog")
	beego.Router("/main/blog/addblog.html", &admin.BlogController{}, "post:AddBlogDo")
	beego.Router("/main/blog/editblog.html", &admin.BlogController{}, "get:EditBlog")
	beego.Router("/main/blog/editblog.html", &admin.BlogController{}, "post:EditBlogDo")
	beego.Router("/main/blog/delblog.html", &admin.BlogController{}, "post:DelBlog")
	beego.Router("/main/blog/sortblog.html", &admin.BlogController{}, "post:SortBlog")
	beego.Router("/main/blog/detailblog.html", &admin.BlogController{}, "get:DetailBlog")

	//前台
	beego.Router("/", &home.IndexController{})
}
