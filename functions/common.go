package functions

import (
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

//
// 时间格式化
//
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
