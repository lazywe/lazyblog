package admin

import "github.com/astaxie/beego/orm"

type Option struct {
	Id         int    `orm:"pk;auto;"`            //主键
	Title      string `orm:"default(0);size(32)"` //名称
	Sort       int    `orm:"default(0)"`
	CreateTime uint   `orm:"default(0);size(10)"`
	UpdateTime uint   `orm:"default(0);size(10)"`
}

//
// 查询功能
//
func (this *Option) FindList() (error, []Option) {
	o := orm.NewOrm()
	var option []Option
	_, err := o.QueryTable(this).All(&option)
	if err != nil {
		return err, nil
	}
	return nil, option
}

//
// 添加功能
//
func (this *Option) Add() (error, int) {
	o := orm.NewOrm()
	id, err := o.Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, int(id)
}

//
// 修改功能
//
func (this *Option) Update(id int) (error, int) {
	o := orm.NewOrm()
	this.Id = id
	num, err := o.Update(this)
	if err != nil {
		return err, 0
	}
	return nil, int(num)
}
