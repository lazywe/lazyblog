package admin

type OptionController struct {
	BaseController
}

//
//功能列表
//
func (this *OptionController) Option() {
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

}

//
//修改功能
//
func (this *OptionController) EditOption() {
	this.TplName = "editoption.html"
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
