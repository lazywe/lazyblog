package admin

import (
	"lazyblog/models"
)

type NodeController struct {
	BaseController
}

type NodeList struct {
	Pid    uint
	Id     uint
	Title  string
	Status uint
	Sort   uint
	Url    string
	Ico    string
	Hint   string
	Level  uint
}

var Nodes []NodeList

func (this *NodeController) Node() {
	Nodes = nil
	err, node := new(models.Node).GetNodeList()
	if err == true {
		Nodes = this.Recursive(node, 0, 0)
	}
	this.Data["Lists"] = Nodes
	this.TplName = "main/node/node.html"
}

func (this *NodeController) Recursive(node []models.Node, pid uint, level uint) []NodeList {
	for _, v := range node {
		if v.Pid == pid {
			var tm = NodeList{Pid: v.Pid, Id: v.Id, Ico: v.Ico, Url: v.Url, Sort: v.Sort, Status: v.Status, Level: level * 9, Title: v.Title}
			Nodes = append(Nodes, tm)
			this.Recursive(node, v.Id, (level + 1))
		}
	}
	return Nodes
}
