package home

import (
	"lazyblog/models"
	"strconv"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Article() {
	id := this.Input().Get("id")
	intid, err := strconv.Atoi(id)
	if intid < 0 || err != nil {
		this.Redirect("/", 302)
	}

	//查出单条博文
	modelblog := new(models.Blog)
	err, result := modelblog.GetBlogInfo(intid)
	if err != nil {
		this.Redirect("/", 302)
	}

	//新增点击数量
	err2, _ := modelblog.UpdateHitsBlog(intid)
	if err2 != nil {
		this.Redirect("/", 302)
	}
	this.Data["val"] = result
	this.TplName = "home/article/article.html"
}
