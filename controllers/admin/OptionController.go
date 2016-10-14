package admin

import (
	"lazyblog/models/admin"
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
	var optionMode = new(admin.Option)
	_, option := optionMode.FindList()
	this.Data["Lists"] = option
	this.TplName = "option/option.html"
	return
}

//
//添加功能
//
func (this *OptionController) AddOption() {
	this.TplName = "option/addoption.html"
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
	optionMode := &admin.Option{Title: title, Sort: sortid, CreateTime: this.Time}
	err, _ := optionMode.Add()
	if err == nil {
		this.AjaxReturn("1", "添加成功", nil)
	} else {
		this.AjaxReturn("0", "添加失败", nil)
	}
}

//
//修改功能
//
func (this *OptionController) EditOption() {
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
	optionMode := &admin.Option{Title: title, Sort: sortid, UpdateTime: this.Time}
	autoid, _ := strconv.Atoi(id)
	err, _ := optionMode.Update(autoid)
	if err == nil {
		this.AjaxReturn("1", "修改成功", nil)
	} else {
		this.AjaxReturn("0", "修改失败", nil)
	}
	return
}

//
//修改功能
//
func (this *OptionController) EditOptionDo() {

}

//
//删除功能
//
func (this *OptionController) DelOption() {

}
