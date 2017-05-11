package home

import "lazyblog/models"

type IndexController struct {
	BaseController
}

func (i *IndexController) Get() {
	var articlemodel = new(models.Blog)
	// 获取总数量
	_, count := articlemodel.GetHomeBlogCount()
	page := i.PageUtil(count, i.PageSize)
	_, list := articlemodel.GetHomeBlogList(page.Offset, i.PageSize)
	i.Data["Blogs"] = list
	i.TplName = "home/index/index.html"
}
