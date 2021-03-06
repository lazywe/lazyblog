package admin

import (
	"lazyblog/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

type BlogController struct {
	BaseController
}

//
//功能列表
//
func (this *BlogController) Blog() {
	var blogMode = new(models.Blog)
	_, blog := blogMode.GetBlogList(nil)
	this.Data["Lists"] = blog
	this.setTplName("blog/blog")
	return
}

//
//添加功能
//
func (this *BlogController) AddBlog() {

	//博客分类
	var categoryMode = new(models.Category)
	_, category := categoryMode.GetCategoryList()
	this.Data["Category"] = category

	//推荐
	var recommendMode = new(models.Recommend)
	_, recommend := recommendMode.GetAdminRecommendList()
	this.Data["Recommend"] = recommend

	this.setTplName("blog/addblog")
	return
}

//
//添加功能
//
func (this *BlogController) AddBlogDo() {
	title := this.Input().Get("title")
	category_id := this.Input().Get("category_id")
	sort := this.Input().Get("sort")
	content := this.Input().Get("content")
	description := this.Input().Get("description")
	state := this.Input().Get("state")
	recommends := make([]string, 0)
	this.Ctx.Input.Bind(&recommends, "recommend")

	valid := validation.Validation{}
	valid.Required(title, "title").Message("请输入标题")
	valid.Required(category_id, "category_id").Message("请选择分类")
	valid.Required(content, "content").Message("请输入博文")
	valid.Required(state, "state").Message("请选择状态")
	if valid.HasErrors() {
		// 打印错误信息
		for _, err := range valid.Errors {
			this.AjaxReturn("0", err.Message, nil)
			return
		}
	}

	sortid, _ := strconv.Atoi(sort)
	stateid, _ := strconv.Atoi(state)
	Category := new(models.Category)
	intcategory_id, _ := strconv.Atoi(category_id)
	Recommend := strings.Join(recommends, ",")

	Category.Id = intcategory_id
	blogMode := &models.Blog{Title: title, Sort: sortid, Description: description, Content: content, Category: Category, State: uint(stateid), CreateTime: this.Time, Recommend: Recommend}
	err, _ := blogMode.AddBlog()
	if err == nil {
		this.AjaxReturn("1", "添加成功", nil)
		return
	} else {
		this.AjaxReturn("0", "添加失败", nil)
		return
	}
}

//
//修改功能
//
func (this *BlogController) EditBlog() {
	id := this.Input().Get("id")
	idint, _ := strconv.Atoi(id)
	var blogmodel = new(models.Blog)
	err, result := blogmodel.GetBlogInfo(idint)
	if err != nil {
		this.Redirect(this.URLFor("AdminController.Main"), 302)
		return
	}

	//选择推荐
	var recommendMode = new(models.Recommend)
	_, recommend := recommendMode.GetAdminRecommendList()
	this.Data["Recommend"] = recommend

	//博客分类
	var categoryMode = new(models.Category)
	_, category := categoryMode.GetCategoryList()
	this.Data["Category"] = category
	this.Data["Val"] = result
	this.setTplName("blog/editblog")
	return
}

//
//修改功能
//
func (this *BlogController) EditBlogDo() {
	title := this.Input().Get("title")
	category_id := this.Input().Get("category_id")
	sort := this.Input().Get("sort")
	content := this.Input().Get("content")
	description := this.Input().Get("description")
	state := this.Input().Get("state")
	id := this.Input().Get("id")
	recommends := make([]string, 0)
	this.Ctx.Input.Bind(&recommends, "recommend")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("非法操作")
	valid.Required(title, "title").Message("请输入标题")
	valid.Required(category_id, "category_id").Message("请选择分类")
	valid.Required(content, "content").Message("请输入博文")
	valid.Required(state, "state").Message("请选择状态")

	if valid.HasErrors() {
		// 打印错误信息
		for _, err := range valid.Errors {
			this.AjaxReturn("0", err.Message, nil)
			return
		}
	}

	sortid, _ := strconv.Atoi(sort)
	autoid, _ := strconv.Atoi(id)
	stateid, _ := strconv.Atoi(state)
	Category := new(models.Category)
	intcategory_id, _ := strconv.Atoi(category_id)
	Recommend := strings.Join(recommends, ",")
	Category.Id = intcategory_id
	blogMode := &models.Blog{Title: title, Sort: sortid, Description: description, Content: content, Category: Category, State: uint(stateid), UpdateTime: this.Time, Recommend: Recommend}
	err, _ := blogMode.UpdateBlog(autoid)
	if err == nil {
		this.AjaxReturn("1", "修改成功", nil)
		return
	} else {
		this.AjaxReturn("0", "修改失败", nil)
		return
	}
}

//
//删除功能
//
func (this *BlogController) DelBlog() {
	id := this.Input().Get("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("非法操作")
	if valid.HasErrors() {
		// 打印错误信息
		for _, err := range valid.Errors {
			this.AjaxReturn("0", err.Message, nil)
			return
		}
	}
	autoid, _ := strconv.Atoi(id)
	var blogmodel = new(models.Blog)
	err, _ := blogmodel.DelBlog(autoid)
	if err == nil {
		this.AjaxReturn("1", "删除成功", nil)
		return
	} else {
		this.AjaxReturn("0", "删除失败", nil)
		return
	}
}

//
// 查看详情
//
func (this *BlogController) DetailBlog() {
	id := this.Input().Get("id")
	idint, _ := strconv.Atoi(id)
	var blogmodel = new(models.Blog)
	err, result := blogmodel.GetBlogInfo(idint)
	if err != nil {
		this.Redirect(this.URLFor("AdminController.Main"), 302)
		return
	}
	this.Data["Val"] = result
	this.setTplName("blog/detailblog")
	return
}

//
//更新排序
//
func (this *BlogController) SortBlog() {
	ids := make([]string, 0)
	sorts := make([]string, 0)
	this.Ctx.Input.Bind(&ids, "id")
	this.Ctx.Input.Bind(&sorts, "sort")
	if ids == nil {
		this.AjaxReturn("0", "更新失败", nil)
		return
	}
	var blogmodel = new(models.Blog)
	for k, v := range ids {
		sortid, _ := strconv.Atoi(sorts[k])
		id, _ := strconv.Atoi(v)
		blogmodel.Sort = sortid
		err, _ := blogmodel.SortBlog(id)
		if err != nil {
			this.AjaxReturn("0", "更新失败", nil)
			return
		}
	}
	this.AjaxReturn("1", "更新成功", nil)
	return
}
