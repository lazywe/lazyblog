package admin

import (
	"lazyblog/models"
	"strconv"

	"github.com/astaxie/beego/validation"
)

type CategoryController struct {
	BaseController
}

//
//功能列表
//
func (this *CategoryController) Category() {
	var categoryMode = new(models.Category)
	_, category := categoryMode.GetCategoryList()
	this.Data["Lists"] = category
	this.setTplName("category/category")
	return
}

//
//添加功能
//
func (this *CategoryController) AddCategory() {
	this.setTplName("category/addcategory")
	return
}

//
//添加功能
//
func (this *CategoryController) AddCategoryDo() {
	title := this.Input().Get("title")
	sort := this.Input().Get("sort")
	state := this.Input().Get("state")

	valid := validation.Validation{}
	valid.Required(title, "title").Message("请输入标题")
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
	categoryMode := &models.Category{Title: title, Sort: sortid, State: uint(stateid), CreateTime: this.Time}
	err, _ := categoryMode.AddCategory()
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
func (this *CategoryController) EditCategory() {
	id := this.Input().Get("id")
	idint, _ := strconv.Atoi(id)
	var categorymodel = new(models.Category)
	err, result := categorymodel.GetCategoryInfo(idint)
	if err != nil {
		this.Redirect(this.URLFor("AdminController.Main"), 302)
		return
	}
	this.Data["Val"] = result
	this.setTplName("category/editcategory")
	return
}

//
//修改功能
//
func (this *CategoryController) EditCategoryDo() {
	title := this.Input().Get("title")
	sort := this.Input().Get("sort")
	state := this.Input().Get("state")
	id := this.Input().Get("id")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("非法操作")
	valid.Required(title, "title").Message("请输入标题")
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
	categoryMode := &models.Category{Title: title, Sort: sortid, State: uint(stateid), UpdateTime: this.Time}
	err, _ := categoryMode.UpdateCategory(autoid)
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
func (this *CategoryController) DelCategory() {
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
	var categorymodel = new(models.Category)
	err, _ := categorymodel.DelCategory(autoid)
	if err == nil {
		this.AjaxReturn("1", "删除成功", nil)
		return
	} else {
		this.AjaxReturn("0", "删除失败", nil)
		return
	}
}

//
//更新排序
//
func (this *CategoryController) SortCategory() {
	ids := make([]string, 0)
	sorts := make([]string, 0)
	this.Ctx.Input.Bind(&ids, "id")
	this.Ctx.Input.Bind(&sorts, "sort")
	if ids == nil {
		this.AjaxReturn("0", "更新失败", nil)
		return
	}
	var categorymodel = new(models.Category)
	for k, v := range ids {
		sortid, _ := strconv.Atoi(sorts[k])
		id, _ := strconv.Atoi(v)
		categorymodel.Sort = sortid
		err, _ := categorymodel.SortCategory(id)
		if err != nil {
			this.AjaxReturn("0", "更新失败", nil)
			return
		}
	}
	this.AjaxReturn("1", "更新成功", nil)
	return
}
