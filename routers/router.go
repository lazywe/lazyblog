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

	//功能类型
	beego.Router("/main/option/option.html", &admin.OptionController{}, "get:Option")
	beego.Router("/main/option/addoption.html", &admin.OptionController{}, "get:AddOption")
	beego.Router("/main/option/addoption.html", &admin.OptionController{}, "post:AddOptionDo")
	beego.Router("/main/option/editoption.html", &admin.OptionController{}, "get:EditOption")
	beego.Router("/main/option/editoption.html", &admin.OptionController{}, "post:EditOptionDo")
	beego.Router("/main/option/deloption.html", &admin.OptionController{}, "post:DelOption")
	beego.Router("/main/option/sortoption.html", &admin.OptionController{}, "post:SortOption")

	//博客管理
	beego.Router("/main/blog/blog.html", &admin.BlogController{}, "get:Blog")
	beego.Router("/main/blog/addblog.html", &admin.BlogController{}, "get:AddBlog")
	beego.Router("/main/blog/addblog.html", &admin.BlogController{}, "post:AddBlogDo")
	beego.Router("/main/blog/editblog.html", &admin.BlogController{}, "get:EditBlog")
	beego.Router("/main/blog/editblog.html", &admin.BlogController{}, "post:EditBlogDo")
	beego.Router("/main/blog/delblog.html", &admin.BlogController{}, "post:DelBlog")
	beego.Router("/main/blog/sortblog.html", &admin.BlogController{}, "post:SortBlog")
	beego.Router("/main/blog/detailblog.html", &admin.BlogController{}, "get:DetailBlog")

	//分类管理
	beego.Router("/main/category/category.html", &admin.CategoryController{}, "get:Category")
	beego.Router("/main/category/addcategory.html", &admin.CategoryController{}, "get:AddCategory")
	beego.Router("/main/category/addcategory.html", &admin.CategoryController{}, "post:AddCategoryDo")
	beego.Router("/main/category/editcategory.html", &admin.CategoryController{}, "get:EditCategory")
	beego.Router("/main/category/editcategory.html", &admin.CategoryController{}, "post:EditCategoryDo")
	beego.Router("/main/category/delcategory.html", &admin.CategoryController{}, "post:DelCategory")
	beego.Router("/main/category/sortcategory.html", &admin.CategoryController{}, "post:SortCategory")

	//栏目管理
	beego.Router("/main/menu/menu.html", &admin.MenuController{}, "get:Menu")
	beego.Router("/main/menu/addmenu.html", &admin.MenuController{}, "get:AddMenu")
	beego.Router("/main/menu/addmenu.html", &admin.MenuController{}, "post:AddMenuDo")
	beego.Router("/main/menu/editmenu.html", &admin.MenuController{}, "get:EditMenu")
	beego.Router("/main/menu/editmenu.html", &admin.MenuController{}, "post:EditMenuDo")
	beego.Router("/main/menu/delmenu.html", &admin.MenuController{}, "post:DelMenu")
	beego.Router("/main/menu/sortmenu.html", &admin.MenuController{}, "post:SortMenu")

	//节点管理
	beego.Router("/main/node/node.html", &admin.NodeController{}, "get:Node")
	beego.Router("/main/node/addnode.html", &admin.NodeController{}, "get:AddNode")
	beego.Router("/main/node/addnode.html", &admin.NodeController{}, "post:AddNodeDo")
	beego.Router("/main/node/editnode.html", &admin.NodeController{}, "get:EditNode")
	beego.Router("/main/node/editnode.html", &admin.NodeController{}, "post:EditNodeDo")
	beego.Router("/main/node/delnode.html", &admin.NodeController{}, "post:DelNode")
	beego.Router("/main/node/sortnode.html", &admin.NodeController{}, "post:SortNode")

	//前台
	beego.Router("/", &home.IndexController{})
	beego.Router("/article.html", &home.ArticleController{}, "get:Article")
}
