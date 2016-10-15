package admin

import (
	"lazyblog/models/admin"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

type Menu struct {
	Title string
	Url   string
	Ico   string
	Id    uint
	Son   []*Menu
}

func (this *AdminController) Main() {
	this.Data["title"] = "我的博客"
	var node = new(admin.Node)
	err, nodelist := node.GetNodeList()
	// err, node := admin.Node.GetNodeList()
	var Menus []*Menu
	if err != true {
		Menus = []*Menu{
			&Menu{Title: "网站配置", Url: this.URLFor("AdminController.Panal"), Ico: "fa-close", Son: []*Menu{
				&Menu{Title: "网站配置", Url: this.URLFor("AdminController.Panal"), Ico: "fa-close"},
			}},
		}
	} else {
		Menus = this.MenuList(nodelist, 0)
	}
	this.Data["Menus"] = Menus
	this.TplName = "main/index/index.html"
}

//递归节点
func (this *AdminController) MenuList(node []admin.Node, pid uint) []*Menu {
	var M []*Menu
	for _, v := range node {
		if v.Pid == pid {
			tm := new(Menu)
			tm.Title = v.Title
			tm.Url = v.Url
			tm.Id = v.Id
			tm.Ico = v.Ico
			M = append(M, tm)
			tm.Son = this.MenuList(node, v.Id)
		}
	}
	return M
}

func (this *AdminController) Panal() {
	this.TplName = "main/index/panal.html"
}
