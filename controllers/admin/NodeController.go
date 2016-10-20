package admin

import (
	"lazyblog/models"
	"strconv"

	"github.com/astaxie/beego/validation"
)

type NodeController struct {
	BaseController
}

type NodeList struct {
	Pid   int
	Id    int
	Title string
	State uint
	Sort  int
	Url   string
	Ico   string
	Hint  string
	Level uint
}

var Nodes []NodeList

//
// 节点列表
//
func (this *NodeController) Node() {
	Nodes = nil
	err, node := new(models.Node).GetNodeList()
	if err == true {
		this.Recursive(node, 0, 0)
	}
	this.Data["Lists"] = Nodes
	this.TplName = "node/node.html"
}

func (this *NodeController) Recursive(node []models.Node, pid int, level uint) {
	for _, v := range node {
		if v.Pid == pid {
			var tm = NodeList{Pid: v.Pid, Id: v.Id, Ico: v.Ico, Url: v.Url, Sort: v.Sort, State: v.State, Level: level * 9, Title: v.Title}
			Nodes = append(Nodes, tm)
			this.Recursive(node, v.Id, (level + 1))
		}
	}
}

//
// 添加节点
//
func (this *NodeController) AddNode() {
	Nodes = nil
	err, node := new(models.Node).GetNodeList()
	if err == true {
		this.Recursive(node, 0, 0)
	}
	this.Data["Nodes"] = Nodes
	this.TplName = "node/addnode.html"
	return
}

//
// 添加
//
func (this *NodeController) AddNodeDo() {
	title := this.Input().Get("title")
	pid := this.Input().Get("pid")
	url := this.Input().Get("url")
	position := this.Input().Get("position")
	hint := this.Input().Get("hint")
	ico := this.Input().Get("ico")
	sort := this.Input().Get("sort")
	state := this.Input().Get("state")

	valid := validation.Validation{}
	valid.Required(title, "title").Message("请输入标题")
	valid.Required(pid, "pid").Message("请选择上级节点")
	valid.Required(url, "url").Message("请输入节点地址")
	valid.Required(position, "position").Message("请选择功能位置")
	valid.Required(hint, "hint").Message("请选择提示类型")
	valid.Required(state, "state").Message("请选择状态")
	if valid.HasErrors() {
		// 打印错误信息
		for _, err := range valid.Errors {
			this.AjaxReturn("0", err.Message, nil)
			return
		}
	}

	idsort, _ := strconv.Atoi(sort)
	intstate, _ := strconv.Atoi(state)
	intpid, _ := strconv.Atoi(pid)
	intposition, _ := strconv.Atoi(position)
	inthint, _ := strconv.Atoi(hint)
	nodeMode := &models.Node{Title: title, Pid: intpid, Url: url, Position: uint(intposition), Hint: uint(inthint), Ico: ico, Sort: idsort, State: uint(intstate)}
	err, _ := nodeMode.AddNode()
	if err == nil {
		this.AjaxReturn("1", "添加成功", nil)
		return
	} else {
		this.AjaxReturn("0", "添加失败", nil)
		return
	}
}

//
//修改功能
//
func (this *NodeController) EditNode() {
	id := this.Input().Get("id")
	idint, _ := strconv.Atoi(id)
	var nodemodel = new(models.Node)
	err, result := nodemodel.GetNodeInfo(idint)
	if err != nil {
		this.Redirect(this.URLFor("AdminController.Main"), 302)
		return
	}

	Nodes = nil
	err2, node := nodemodel.GetNodeList()
	if err2 == true {
		this.Recursive(node, 0, 0)
	}
	this.Data["Nodes"] = Nodes
	this.Data["Val"] = result
	this.Data["Hint"] = "1"
	this.TplName = "node/editnode.html"
	return
}

func (this *NodeController) EditNodeDo() {

}

//
//删除
//
func (this *NodeController) DelNode() {
	// id := this.Input().Get("id")
	// valid := validation.Validation{}
	// valid.Required(id, "id").Message("非法操作")
	// if valid.HasErrors() {
	// 	// 打印错误信息
	// 	for _, err := range valid.Errors {
	// 		this.AjaxReturn("0", err.Message, nil)
	// 		return
	// 	}
	// }
	// autoid, _ := strconv.Atoi(id)
	// var nodemodel = new(models.Node)
	// err, _ := nodemodel.DelNode(autoid)
	// if err == nil {
	// 	this.AjaxReturn("1", "删除成功", nil)
	// 	return
	// } else {
	// 	this.AjaxReturn("0", "删除失败", nil)
	// 	return
	// }
}

//
//更新排序
//
func (this *NodeController) SortNode() {
	// ids := make([]string, 0)
	// sorts := make([]string, 0)
	// this.Ctx.Input.Bind(&ids, "id")
	// this.Ctx.Input.Bind(&sorts, "sort")
	// if ids == nil {
	// 	this.AjaxReturn("0", "更新失败", nil)
	// 	return
	// }
	// var nodemodel = new(models.Node)
	// for k, v := range ids {
	// 	sortid, _ := strconv.Atoi(sorts[k])
	// 	id, _ := strconv.Atoi(v)
	// 	nodemodel.Sort = sortid
	// 	err, _ := nodemodel.SortNode(id)
	// 	if err != nil {
	// 		this.AjaxReturn("0", "更新失败", nil)
	// 		return
	// 	}
	// }
	// this.AjaxReturn("1", "更新成功", nil)
	// return
}
