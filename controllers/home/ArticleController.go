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

	modleblog := new(models.Blog)
	err, result := modleblog.GetBlogInfo(intid)
	if err != nil {
		this.Redirect("/", 302)
	}
	this.Data["val"] = result
	this.TplName = "article/article.html"
}
