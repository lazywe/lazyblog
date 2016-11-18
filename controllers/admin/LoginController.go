package admin

import (
	"lazyblog/models"

	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	BaseController
}

//
// 登录
//
func (this *LoginController) Login() {
	this.TplName = "admin/login.html"
}

//
// 处理登录
//
func (this *LoginController) LoginDo() {

	user := this.Input().Get("user")
	pass := this.Input().Get("pass")

	valid := validation.Validation{}
	valid.Required(user, "user").Message("请输入用户名或者密码")
	valid.Required(pass, "pass").Message("请输入用户名或者密码")
	if valid.HasErrors() {
		// 打印错误信息
		for _, err := range valid.Errors {
			this.AjaxReturn("0", err.Message, nil)
			return
		}
	}

	err, result := models.GetUserInfo(user)
	if err == false {
		this.AjaxReturn("0", "登录失败，用户名或密码错误", nil)
		return
	}

	hspass := this.md5(pass, result.Salt)
	if hspass != result.Pass {
		this.AjaxReturn("0", "登录失败密码错误", nil)
		return
	}
	//设置登录状态
	auth := map[string]interface{}{"user": result.User}
	this.SetAuth(auth)
	//返回登录状态
	data := map[string]interface{}{"url": this.URLFor("AdminController.Main")}

	// data := map[string]interface{}{"url": this.URLFor("BlogController.Blog")}
	// data := map[string]interface{}{"url": this.URLFor("NodeController.Node")}
	this.AjaxReturn("1", "登录成功", data)
	return
}
