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
	this.getMenu()
	this.getRecommend()
}

//
// 导航
//
func (this *BaseController) getMenu() {

	if Menus == nil {
		//导航
		var menuMode = new(models.Menu)
		_, Menus = menuMode.GetHomeMenuList()
	}
	this.Data["Menus"] = Menus
}

type RecommendsData struct {
	Title string
	Blog  []models.Blog
}

//
// 推荐栏目
//
func (this *BaseController) getRecommend() {
	//栏目
	var RecommendMode = new(models.Recommend)
	err, Recommends := RecommendMode.GetHomeRecommendList()
	var RecommendsDatas []*RecommendsData
	var BlogModel = new(models.Blog)
	//不同类型获取方法不一
	if err == nil {
		for _, v := range Recommends {
			if v.Id == 2 { //热门博文
				err, recommendBlogs := BlogModel.GetRmBlogList()
				if err == nil {
					RecommendsDatas = append(RecommendsDatas, &RecommendsData{Title: v.Title, Blog: recommendBlogs})
				}
			} else { //精品博文
				err, recommendBlogs := BlogModel.GetTjBlogList(v.Id)
				if err == nil {
					RecommendsDatas = append(RecommendsDatas, &RecommendsData{Title: v.Title, Blog: recommendBlogs})
				}
			}
		}
		this.Data["RecommendsDatas"] = RecommendsDatas
	}
}
