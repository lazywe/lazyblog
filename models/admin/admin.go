package admin

import "github.com/astaxie/beego/orm"

type Admin struct {
	Id   int    `orm:"pk;auto;` //主键
	User string `orm:"default(0);size(32)"`
	Pass string `orm:"default(0);size(32)"`
	Salt string `orm:"default(0);size(32)"`
}

func GetUserInfo(user string) (bool, Admin) {
	o := orm.NewOrm()
	var admin Admin
	err := o.QueryTable(admin).Filter("user", user).One(&admin)
	return err != orm.ErrNoRows, admin
}
