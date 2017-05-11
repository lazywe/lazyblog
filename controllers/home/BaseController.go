package home

import (
	"fmt"
	"lazyblog/models"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	PageSize int
	Time     uint
}

var Menus []models.Menu

func (b *BaseController) Prepare() {
	b.Time = uint(time.Now().Unix())
	b.PageSize = 10
	b.getMenu()
	b.getRecommend()
}

/// 分页
type Page struct {
	Offset     int
	PageNum    string
	TotalPage  float64
	TotalCount int64
}

type RecommendsData struct {
	Title string
	Blog  []models.Blog
}

//
// 导航
//
func (b *BaseController) getMenu() {

	if Menus == nil {
		//导航
		var menuMode = new(models.Menu)
		_, Menus = menuMode.GetHomeMenuList()
	}
	b.Data["Menus"] = Menus
}

//
// 推荐栏目
//
func (b *BaseController) getRecommend() {
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
		b.Data["RecommendsDatas"] = RecommendsDatas
	}

}

// 分页
func (b *BaseController) PageUtil(count int64, pagenum int) *Page {
	var currentPage int
	pageGet := b.Input().Get("page")
	totalPage := math.Ceil(float64(count) / float64(pagenum))
	if pageGet == "" {
		currentPage = 1
	} else {
		idpage, _ := strconv.Atoi(pageGet)
		currentPage = idpage
		// 若大于最后一页则为最后一页
		if float64(currentPage) > totalPage {
			currentPage = int(totalPage)
			// 若小于0则为1
		} else if currentPage < 0 {
			currentPage = 1
		}
	}
	offset := (currentPage - 1) * pagenum
	pageData := map[string]interface{}{
		// 上一页
		"UpPageUrl": "",
		/// 下一页
		"DownPageUrl": "",
		/// 当前页码
		"CurrentPage": currentPage,
		/// 总条数
		"Count": count,
	}

	var NumPage = []map[string]interface{}{}

	// 当前地址
	currentUrl := b.Ctx.Request.URL.Path

	/// 所有参数
	currentQuery := b.Ctx.Request.URL.RawQuery
	for i := 1; i <= int(totalPage); i++ {
		var curmap = map[string]interface{}{}
		var currentPageUrl string = ""

		// 页码
		if pageGet == "" {
			currentPageUrl = currentUrl + "?page=" + strconv.Itoa(i) + currentQuery
		} else {
			currentPageUrl = currentUrl + "?" + strings.Replace(currentQuery, "page="+pageGet, "page="+strconv.Itoa(i), -1)
		}
		// 上一页
		if currentPage-1 == i && i-1 >= 0 {
			pageData["UpPageUrl"] = currentPageUrl
		}

		// 下一页
		if currentPage+1 == i && i <= int(totalPage) {
			pageData["DownPageUrl"] = currentPageUrl
		}

		// 是否当前页
		if currentPage == i {
			curmap["isCurrentPage"] = true
		} else {
			curmap["isCurrentPage"] = false
		}
		curmap["url"] = currentPageUrl
		curmap["currentpage"] = strconv.Itoa(i)
		NumPage = append(NumPage, curmap)
	}
	fmt.Println(pageData)
	pageData["NumPage"] = NumPage
	b.Data["Page"] = pageData
	return &Page{Offset: offset, PageNum: strconv.Itoa(pagenum), TotalCount: count, TotalPage: totalPage}
}
