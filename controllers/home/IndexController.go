package home

import "lazyblog/models"

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {

	var articlemodel = new(models.Blog)
	_, list := articlemodel.GetHomeBlogList()
	this.Data["Blogs"] = list
	this.TplName = "index/index.html"
}
