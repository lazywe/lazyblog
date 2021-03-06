package functions

import (
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

//
// 初始化函数
//
func Init() {
	beego.AddFuncMap("str_replace", str_replace)
	beego.AddFuncMap("time_format", time_format)
	beego.AddFuncMap("checked", checked)
	beego.AddFuncMap("selected", checked)
	beego.AddFuncMap("recommendCheck", recommendCheck)
}

//
// 字符串替换函数，根据数量替换
//
// @param  uint    in   需要替换的数量
// @param  string  str  需要替换的字符串
// @return string  out  替换后的结果
func str_replace(in uint, str string) (out string) {
	return strings.Repeat(str, int(in))
}

//
// 时间格式化
//
// @param uint   in     需要替换的事件戳
// @param string format 替换的格式  Y-m-d H:i:s
func time_format(in uint, format string) (out string) {
	if in == 0 {
		return "null"
	}
	format = strings.Replace(format, "Y", "2006", 5)
	format = strings.Replace(format, "y", "2006", 5)
	format = strings.Replace(format, "m", "01", 5)
	format = strings.Replace(format, "d", "02", 5)
	format = strings.Replace(format, "H", "15", 5)
	format = strings.Replace(format, "h", "15", 5)
	format = strings.Replace(format, "i", "04", 5)
	format = strings.Replace(format, "s", "04", 5)
	t := time.Unix(int64(in), 0)
	return t.Format(format)
}

//
// checked 比较函数
//
func checked(arg1 interface{}, arg2 interface{}) bool {
	switch arg1.(type) {
	case uint:
		newarg1 := int(arg1.(uint))
		if newarg1 != arg2.(int) {
			return false
		}
	case string:
		newarg1 := arg1.(string)
		if newarg1 != arg2.(string) {
			return false
		}
	default:
		newarg1 := arg1
		if newarg1 != arg2 {
			return false
		}
	}
	return true
}

//
// 推荐判断
//
func recommendCheck(id int, recs string) string {
	tempArr := strings.Split(recs, ",")
	for _, v := range tempArr {
		intV, _ := strconv.Atoi(v)
		if id == intV {
			return "checked"
			break
		}
	}
	return ""
}
