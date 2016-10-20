package models

import "github.com/astaxie/beego/orm"

type Menu struct {
	Id         int    `orm:"pk;auto;"`
	Title      string `orm:"default(0);size(32);"`
	Link       string `orm:"default(0);size(128);"`
	Sort       int    `orm:"default(0);`
	CreateTime uint   `orm:size(10)`
	State      uint   `orm:"default(0);size(1)"`
	UpdateTime uint   `orm:"default(0);size(10)"`
}

//
// 查询功能列表
//
func (this *Menu) GetMenuList() (error, []Menu) {
	o := orm.NewOrm()
	var blog []Menu
	result := o.QueryTable(this)
	result = result.OrderBy("-Sort")
	_, err := result.All(&blog)
	if err != nil {
		return err, nil
	}
	return nil, blog
}

//
// 添加功能
//
func (this *Menu) AddMenu() (error, int) {
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
func (this *Menu) UpdateMenu(id int) (error, int) {
	o := orm.NewOrm()
	this.Id = id
	num, err := o.Update(this, "Title", "Sort", "State", "Link", "UpdateTime")
	if err != nil {
		return orm.ErrNoRows, 0
	}
	return nil, int(num)
}

//
// 读取功能
//
func (this *Menu) GetMenuInfo(id int) (error, *Menu) {
	o := orm.NewOrm()
	this.Id = id
	err := o.Read(this)
	if err != nil {
		return err, nil
	}
	return nil, this
}

//
// 删除
//
func (this *Menu) DelMenu(id int) (error, int) {
	o := orm.NewOrm()
	this.Id = id
	num, err := o.Delete(this)
	if err != nil {
		return err, 0
	}
	return nil, int(num)
}

//
// 排序
//
func (this *Menu) SortMenu(id int) (error, int) {
	o := orm.NewOrm()
	this.Id = id
	num, err := o.Update(this, "Sort")
	if err != nil {
		return err, 0
	}
	return nil, int(num)
}

/***************************************** 前台查询 *********************************************/
/***************************************** 前台查询 *********************************************/
/***************************************** 前台查询 *********************************************/
/***************************************** 前台查询 *********************************************/

//
// 查询功能列表
//
func (this *Menu) GetHomeMenuList() (error, []Menu) {
	o := orm.NewOrm()
	var blog []Menu
	result := o.QueryTable(this)
	result = result.OrderBy("-Sort").Filter("State", 1)
	_, err := result.All(&blog, "Title", "Link")
	if err != nil {
		return err, nil
	}
	return nil, blog
}
