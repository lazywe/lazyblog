package models

import "github.com/astaxie/beego/orm"

type Recommend struct {
	Id      int    `orm:"pk;auto;"`
	Title   string `orm:"default(0);size(32)"` //名称
	State   uint   `orm:"default(1);size(1)"`  // 是否前台显示  1显示  0隐藏
	IsAdmin int    `orm:"default(1);size(1)"`  // 是否后台添加 1是 0否
	Sort    int    `orm:"default(0);`
}

//
// 获取后台comment
//
func (this *Recommend) GetAdminRecommendList() (error, []Recommend) {
	o := orm.NewOrm()
	var Recommends []Recommend
	result := o.QueryTable(this)
	result = result.Filter("IsAdmin", 1)
	result = result.OrderBy("-Sort")
	_, err := result.All(&Recommends)
	if err != nil {
		return err, nil
	}
	return nil, Recommends
}

//
// 获取前台comment
//
func (this *Recommend) GetHomeRecommendList() (error, []Recommend) {
	o := orm.NewOrm()
	var Recommends []Recommend
	result := o.QueryTable(this)
	result = result.Filter("State", 1)
	result = result.OrderBy("-Sort")
	_, err := result.All(&Recommends)
	if err != nil {
		return err, nil
	}
	return nil, Recommends
}
