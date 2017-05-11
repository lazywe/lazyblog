package admin

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Time       uint
	Auth       map[string]interface{}
	SessionKey string
}

//
// 初始化函数
//
func (this *BaseController) Prepare() {
	// beego.SetViewsPath(beego.AppConfig.String("admin.tpl"))
	this.Time = uint(time.Now().Unix())
	this.SessionKey = "Admin"
	c, _ := this.GetControllerAndAction()
	upperc := strings.ToLower(c)
	if upperc != "logincontroller" {
		islogin, auth := this.GetAuth()
		if islogin == false {
			this.Redirect(this.URLFor("LoginController.Login"), 302)
			return
		}
		this.Auth = auth.(map[string]interface{})
		log.Println("当前用户")
		log.Println(this.Auth["user"])
		log.Println(auth)
	}
}

// 设置后台模版路径，与模版主题
func (this *BaseController) setTplName(file string) {
	this.TplName = beego.AppConfig.String("admin.tpl") + file + beego.AppConfig.String("tpl.ext")
}

//
// Ajax返回信息
//
func (this *BaseController) AjaxReturn(status string, info string, data interface{}) {
	this.Data["json"] = map[string]interface{}{"status": status, "info": info, "data": data}
	this.ServeJSON()
	return
}

//
// 判断用户是否登录
// @return 登录状态，登录数据
//
func (this *BaseController) GetAuth() (bool, interface{}) {
	auth := this.GetSession(this.SessionKey)
	if auth == nil {
		return false, nil
	}
	return true, auth
}

//
// 设置登录状态
// @return bool
//
func (this *BaseController) SetAuth(user interface{}) bool {
	this.SetSession(this.SessionKey, user)
	return true
}

//
// 删除登录状态
// @return bool
//
func (this *BaseController) DelAuth(user interface{}) bool {
	this.DelSession(this.SessionKey)
	return true
}

/**
 * 密码验证规则,UCENTER方式
 * @param pass 密码
 */
func (this *BaseController) md5(pass string, salt string) string {
	h := md5.New()
	h.Write([]byte(pass))
	cipherStr := h.Sum(nil)
	h2 := md5.New()
	h2.Write([]byte((hex.EncodeToString(cipherStr) + salt)))
	cipherStrtwo := h2.Sum(nil)
	return hex.EncodeToString(cipherStrtwo)
}
