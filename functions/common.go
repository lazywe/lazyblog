package functions

import (
	"strings"

	"github.com/astaxie/beego"
)

//
// 初始化函数
//
func Init() {
	beego.AddFuncMap("str_replace", str_replace)
}

//
// 字符串替换函数，根据数量替换
//
// @param in uint     需要替换的数量
// @param str string  需要替换的字符串
// @return out string 替换后的结果
func str_replace(in uint, str string) (out string) {
	return strings.Repeat(str, int(in))
}
