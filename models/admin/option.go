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
// 查询功能列表
//
func (this *Option) GetOptionList() (error, []Option) {
	o := orm.NewOrm()
	var option []Option
	result := o.QueryTable(this)
	result = result.OrderBy("-Sort")
	_, err := result.All(&option)
	if err != nil {
		return err, nil
	}
	return nil, option
}

//
// 添加功能
//
func (this *Option) AddOption() (error, int) {
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
func (this *Option) UpdateOption(id int) (error, int) {
	o := orm.NewOrm()
	this.Id = id
	num, err := o.Update(this, "Title", "Sort", "UpdateTime")
	if err != nil {
		return orm.ErrNoRows, 0
	}
	return nil, int(num)
}

//
// 读取功能
//
func (this *Option) GetOptionInfo(id int) (error, *Option) {
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
func (this *Option) DelOption(id int) (error, int) {
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
func (this *Option) SortOption(id int) (error, int) {
	o := orm.NewOrm()
	this.Id = id
	num, err := o.Update(this, "Sort")
	if err != nil {
		return err, 0
	}
	return nil, int(num)
}
