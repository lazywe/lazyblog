package admin

import "github.com/astaxie/beego/orm"

type Node struct {
	Id       uint   `orm:"pk;auto"`
	Pid      uint   `orm:"default(0);size(10)"`
	Sort     uint   `orm:"default(999);size(11)"`
	Url      string `orm:"default(0);size(128)"`
	Title    string `orm:"default(0);size(20)"`
	Status   uint   `orm:"default(1);size(1)"`
	Position string `orm:"default('list');size(32)"`
	Ico      string `orm:"size(32)"`
	Hint     string `orm:"default('href');size(32)"`
}

func (this Node) GetNodeList() (bool, []Node) {
	o := orm.NewOrm()
	var nodes []Node
	_, err := o.QueryTable(this).All(&nodes)
	return err != orm.ErrNoRows, nodes
}
