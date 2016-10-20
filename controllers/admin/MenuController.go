package admin

import (
	"lazyblog/models"
	"strconv"

	"github.com/astaxie/beego/validation"
)

type MenuController struct {
	BaseController
}

//
//功能列表
//
func (this *MenuController) Menu() {
	var menuMode = new(models.Menu)
	_, menu := menuMode.GetMenuList()
	this.Data["Lists"] = menu
	this.TplName = "menu/menu.html"
	return
}

//
//添加功能
//
func (this *MenuController) AddMenu() {
	this.TplName = "menu/addmenu.html"
	return
}

//
//添加功能
//
func (this *MenuController) AddMenuDo() {
	title := this.Input().Get("title")
	link := this.Input().Get("link")
	sort := this.Input().Get("sort")
	state := this.Input().Get("state")

	valid := validation.Validation{}
	valid.Required(title, "title").Message("请输入标题")
	valid.Required(link, "link").Message("请输入访问地址")
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
	menuMode := &models.Menu{Title: title, Link: link, Sort: sortid, State: uint(stateid), CreateTime: this.Time}
	err, _ := menuMode.AddMenu()
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
func (this *MenuController) EditMenu() {
	id := this.Input().Get("id")
	idint, _ := strconv.Atoi(id)
	var menumodel = new(models.Menu)
	err, result := menumodel.GetMenuInfo(idint)
	if err != nil {
		this.Redirect(this.URLFor("AdminController.Main"), 302)
		return
	}
	this.Data["Val"] = result
	this.TplName = "menu/editmenu.html"
	return
}

//
//修改功能
//
func (this *MenuController) EditMenuDo() {
	title := this.Input().Get("title")
	link := this.Input().Get("link")
	sort := this.Input().Get("sort")
	state := this.Input().Get("state")
	id := this.Input().Get("id")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("非法操作")
	valid.Required(title, "title").Message("请输入标题")
	valid.Required(link, "link").Message("请输入访问地址")
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
	menuMode := &models.Menu{Title: title, Sort: sortid, Link: link, State: uint(stateid), UpdateTime: this.Time}
	err, _ := menuMode.UpdateMenu(autoid)
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
func (this *MenuController) DelMenu() {
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
	var menumodel = new(models.Menu)
	err, _ := menumodel.DelMenu(autoid)
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
func (this *MenuController) SortMenu() {
	ids := make([]string, 0)
	sorts := make([]string, 0)
	this.Ctx.Input.Bind(&ids, "id")
	this.Ctx.Input.Bind(&sorts, "sort")
	if ids == nil {
		this.AjaxReturn("0", "更新失败", nil)
		return
	}
	var menumodel = new(models.Menu)
	for k, v := range ids {
		sortid, _ := strconv.Atoi(sorts[k])
		id, _ := strconv.Atoi(v)
		menumodel.Sort = sortid
		err, _ := menumodel.SortMenu(id)
		if err != nil {
			this.AjaxReturn("0", "更新失败", nil)
			return
		}
	}
	this.AjaxReturn("1", "更新成功", nil)
	return
}
