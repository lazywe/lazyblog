package admin

import "lazyblog/models"

type AdminController struct {
	BaseController
}

type Menu struct {
	Title string
	Url   string
	Ico   string
	Id    int
	Son   []*Menu
}

func (this *AdminController) Main() {
	this.Data["title"] = "我的博客"
	var node = new(models.Node)
	err, nodelist := node.GetNodeList()
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
	this.setTplName("index/index")
}

//递归节点
func (this *AdminController) MenuList(node []models.Node, pid int) []*Menu {
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
	this.setTplName("index/panal")
}
