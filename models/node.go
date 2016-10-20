package models

import "github.com/astaxie/beego/orm"

type Node struct {
	Id       int    `orm:"pk;auto"`
	Pid      int    `orm:"default(0);"`
	Sort     int    `orm:"default(999);size(11)"`
	Url      string `orm:"default(0);size(128)"`
	Title    string `orm:"default(0);size(20)"`
	State    uint   `orm:"default(1);size(1)"`
	Position uint   `orm:"default(1);size(1)"`
	Ico      string `orm:"size(32)"`
	Hint     uint   `orm:"default(1);size(1)"`
}

//
// 查询
//
func (this *Node) GetNodeList() (bool, []Node) {
	o := orm.NewOrm()
	var nodes []Node
	_, err := o.QueryTable(this).All(&nodes)
	return err != orm.ErrNoRows, nodes
}

//
// 添加功能
//
func (this *Node) AddNode() (error, int) {
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
func (this *Node) UpdateNode(id int) (error, int) {
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
func (this *Node) GetNodeInfo(id int) (error, *Node) {
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
func (this *Node) DelNode(id int) (error, int) {
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
func (this *Node) SortNode(id int) (error, int) {
	o := orm.NewOrm()
	this.Id = id
	num, err := o.Update(this, "Sort")
	if err != nil {
		return err, 0
	}
	return nil, int(num)
}
