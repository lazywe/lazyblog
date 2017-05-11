package admin

import (
	"lazyblog/models"
	"strconv"

	"github.com/astaxie/beego/validation"
)

type OptionController struct {
	BaseController
}

//
//功能列表
//
func (this *OptionController) Option() {
	var optionMode = new(models.Option)
	_, option := optionMode.GetOptionList()
	this.Data["Lists"] = option
	this.setTplName("option/option")
	return
}

//
//添加功能
//
func (this *OptionController) AddOption() {
	this.setTplName("option/addoption")
	return
}

//
//添加功能
//
func (this *OptionController) AddOptionDo() {
	title := this.Input().Get("title")
	sort := this.Input().Get("sort")
	sortid, _ := strconv.Atoi(sort)
	valid := validation.Validation{}
	valid.Required(title, "title").Message("请输入功能名称")
	if valid.HasErrors() {
		// 打印错误信息
		for _, err := range valid.Errors {
			this.AjaxReturn("0", err.Message, nil)
			return
		}
	}
	optionMode := &models.Option{Title: title, Sort: sortid, CreateTime: this.Time}
	err, _ := optionMode.AddOption()
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
func (this *OptionController) EditOption() {
	id := this.Input().Get("id")
	idint, _ := strconv.Atoi(id)
	var optionmodel = new(models.Option)
	err, result := optionmodel.GetOptionInfo(idint)
	if err != nil {
		this.Redirect("/", 1)
	}
	this.Data["Val"] = result
	this.setTplName("option/editoption")
}

//
//修改功能
//
func (this *OptionController) EditOptionDo() {
	title := this.Input().Get("title")
	sort := this.Input().Get("sort")
	sortid, _ := strconv.Atoi(sort)
	id := this.Input().Get("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("非法操作")
	valid.Required(title, "title").Message("请输入功能名称")
	if valid.HasErrors() {
		// 打印错误信息
		for _, err := range valid.Errors {
			this.AjaxReturn("0", err.Message, nil)
			return
		}
	}
	optionMode := &models.Option{Title: title, Sort: sortid, UpdateTime: this.Time}
	autoid, _ := strconv.Atoi(id)
	err, _ := optionMode.UpdateOption(autoid)
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
func (this *OptionController) DelOption() {
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
	var optionmodel = new(models.Option)
	err, _ := optionmodel.DelOption(autoid)
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
func (this *OptionController) SortOption() {
	ids := make([]string, 0)
	sorts := make([]string, 0)
	this.Ctx.Input.Bind(&ids, "id")
	this.Ctx.Input.Bind(&sorts, "sort")
	if ids == nil {
		this.AjaxReturn("0", "更新失败1", nil)
		return
	}
	var optionmodel = new(models.Option)
	for k, v := range ids {
		sortid, _ := strconv.Atoi(sorts[k])
		id, _ := strconv.Atoi(v)
		optionmodel.Sort = sortid
		err, _ := optionmodel.SortOption(id)
		if err != nil {
			this.AjaxReturn("0", "更新失败2", nil)
			return
		}
	}
	this.AjaxReturn("1", "更新成功", nil)
	return
}
