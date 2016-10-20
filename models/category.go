package models

import "github.com/astaxie/beego/orm"

type Category struct {
	Id         int    `orm:"pk;auto;"`
	Title      string `orm:"default(0);size(32);"`
	Sort       int    `orm:"default(0);`
	CreateTime uint   `orm:size(10)`
	State      uint   `orm:"default(0);size(1)"`
	UpdateTime uint   `orm:"default(0);size(10)"`
}

//
// 查询功能列表
//
func (this *Category) GetCategoryList() (error, []Category) {
	o := orm.NewOrm()
	var blog []Category
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
func (this *Category) AddCategory() (error, int) {
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
func (this *Category) UpdateCategory(id int) (error, int) {
	o := orm.NewOrm()
	this.Id = id
	num, err := o.Update(this, "Title", "Sort", "State", "UpdateTime")
	if err != nil {
		return orm.ErrNoRows, 0
	}
	return nil, int(num)
}

//
// 读取功能
//
func (this *Category) GetCategoryInfo(id int) (error, *Category) {
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
func (this *Category) DelCategory(id int) (error, int) {
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
func (this *Category) SortCategory(id int) (error, int) {
	o := orm.NewOrm()
	this.Id = id
	num, err := o.Update(this, "Sort")
	if err != nil {
		return err, 0
	}
	return nil, int(num)
}
