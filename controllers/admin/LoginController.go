package admin

import (
	"crypto/md5"
	"encoding/hex"
	"lazyblog/models/admin"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

var Info string
var Status bool
var Datas string
var Url string = "/"

/**
 * 登录
 */
func (this *LoginController) Login() {
	this.TplName = "main/admin/login.html"
}

/**
 * 登录
 */
func (this *LoginController) LoginDo() {
	user := this.Input().Get("user")
	pass := this.Input().Get("pass")
	if user == "" || pass == "" {
		Status = false
		Info = "请输入用户名或者密码"
	} else {
		err, data := admin.GetUserInfo(user)
		if err == false {
			Status = false
			Info = "登录失败，用户名或密码错误"
		} else {
			hspass := this.md5(pass, data.Salt)
			if hspass != data.Pass {
				Status = false
				Info = "登录失败密码错误"
			} else {
				Status = true
				Info = "登录成功"
				Url = "/main.html"
			}
		}
	}
	this.Data["json"] = map[string]interface{}{"status": Status, "info": Info, "url": Url}
	this.ServeJSON()
	return
}

/**
 * 密码验证规则
 * @param pass 密码
 */
func (this *LoginController) md5(pass string, salt string) string {
	h := md5.New()
	h.Write([]byte(pass))
	cipherStr := h.Sum(nil)
	h2 := md5.New()
	h2.Write([]byte((hex.EncodeToString(cipherStr) + salt)))
	cipherStrtwo := h2.Sum(nil)
	return hex.EncodeToString(cipherStrtwo)
}
