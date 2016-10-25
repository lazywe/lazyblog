package models

import "github.com/astaxie/beego/orm"

type Blog struct {
	Id          int       `orm:"pk;auto"`
	Title       string    `orm:"default(0);size(32)"`
	Content     string    `orm:"type(text)"`
	Sort        int       `orm:"default(0);`
	CreateTime  uint      `orm:"default(0);size(10)"`
	Description string    `orm:"default(0);size(255)"`
	State       uint      `orm:"default(0);size(1)"`
	UpdateTime  uint      `orm:"default(0);size(10)"`
	Category    *Category `orm:"rel(one)"`
}

//
// 查询功能列表
//
func (this *Blog) GetBlogList() (error, []Blog) {
	o := orm.NewOrm()
	var blog []Blog
	// var category Category
	result := o.QueryTable(this).RelatedSel()

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
func (this *Blog) AddBlog() (error, int) {
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
func (this *Blog) UpdateBlog(id int) (error, int) {
	o := orm.NewOrm()
	this.Id = id
	num, err := o.Update(this, "Title", "Sort", "Description", "Content", "State", "UpdateTime")
	if err != nil {
		return orm.ErrNoRows, 0
	}
	return nil, int(num)
}

//
// 读取功能
//
func (this *Blog) GetBlogInfo(id int) (error, *Blog) {
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
func (this *Blog) DelBlog(id int) (error, int) {
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
func (this *Blog) SortBlog(id int) (error, int) {
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

func (this *Blog) GetHomeBlogList() (error, []Blog) {

	o := orm.NewOrm()
	var blog []Blog
	result := o.QueryTable(this)
	result = result.OrderBy("-Sort").Filter("State", 1)
	_, err := result.All(&blog)
	if err != nil {
		return err, nil
	}
	return nil, blog
}
