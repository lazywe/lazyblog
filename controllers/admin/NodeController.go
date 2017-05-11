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
	Pid      int
	Id       int
	Title    string
	Position uint
	State    uint
	Sort     int
	Url      string
	Ico      string
	Hint     uint
	Level    uint
}

var Nodes []NodeList

//
// 列表
//
func (this *NodeController) Node() {
	Nodes = nil
	err, node := new(models.Node).GetNodeList()
	if err == true {
		this.Recursive(node, 0, 0, 0)
	}
	this.Data["Lists"] = Nodes
	this.setTplName("node/node")
}

//
// 计算递归节点数组
//
func (this *NodeController) Recursive(node []models.Node, pid int, level uint, outid int) {
	for _, v := range node {
		if (outid != 0 && outid == v.Pid) || (outid != 0 && outid == v.Id) {
			continue
		}
		if v.Pid == pid {
			var tm = NodeList{Pid: v.Pid, Id: v.Id, Ico: v.Ico, Url: v.Url, Sort: v.Sort, Position: v.Position, Hint: v.Hint, State: v.State, Level: level, Title: v.Title}
			Nodes = append(Nodes, tm)
			this.Recursive(node, v.Id, (level + 1), outid)
		}
	}
}

//
// 添加
//
func (this *NodeController) AddNode() {
	Nodes = nil
	err, node := new(models.Node).GetNodeList()
	if err == true {
		this.Recursive(node, 0, 1, 0)
	}
	this.Data["Nodes"] = Nodes
	this.setTplName("node/addnode")
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
//修改
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
		this.Recursive(node, 0, 1, idint)
	}
	this.Data["Nodes"] = Nodes
	this.Data["Val"] = result
	this.Data["Hint"] = "1"
	this.setTplName("node/editnode")
	return
}

func (this *NodeController) EditNodeDo() {
	id := this.Input().Get("id")
	title := this.Input().Get("title")
	pid := this.Input().Get("pid")
	url := this.Input().Get("url")
	position := this.Input().Get("position")
	hint := this.Input().Get("hint")
	ico := this.Input().Get("ico")
	sort := this.Input().Get("sort")
	state := this.Input().Get("state")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("非法操作")
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

	intid, _ := strconv.Atoi(id)
	idsort, _ := strconv.Atoi(sort)
	intstate, _ := strconv.Atoi(state)
	intpid, _ := strconv.Atoi(pid)
	intposition, _ := strconv.Atoi(position)
	inthint, _ := strconv.Atoi(hint)
	nodeMode := &models.Node{Title: title, Pid: intpid, Url: url, Position: uint(intposition), Hint: uint(inthint), Ico: ico, Sort: idsort, State: uint(intstate)}
	err, _ := nodeMode.UpdateNode(intid)
	if err == nil {
		this.AjaxReturn("1", "修改成功", nil)
		return
	} else {
		this.AjaxReturn("0", "修改失败", nil)
		return
	}
}

//
//删除
//
func (this *NodeController) DelNode() {
	id := this.Input().Get("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("非法操作")
	if valid.HasErrors() {
		// 打印错误信息
		for _, err := range valid.Errors {
			this.AjaxReturn("0", err.Message, nil)
			return
		}
	}
	autoid, _ := strconv.Atoi(id)
	var nodemodel = new(models.Node)
	err, _ := nodemodel.DelNode(autoid)
	if err == nil {
		this.AjaxReturn("1", "删除成功", nil)
		return
	} else {
		this.AjaxReturn("0", "删除失败", nil)
		return
	}
}

//
//更新排序
//
func (this *NodeController) SortNode() {
	ids := make([]string, 0)
	sorts := make([]string, 0)
	this.Ctx.Input.Bind(&ids, "id")
	this.Ctx.Input.Bind(&sorts, "sort")
	if ids == nil {
		this.AjaxReturn("0", "更新失败", nil)
		return
	}
	var nodemodel = new(models.Node)
	for k, v := range ids {
		sortid, _ := strconv.Atoi(sorts[k])
		id, _ := strconv.Atoi(v)
		nodemodel.Sort = sortid
		err, _ := nodemodel.SortNode(id)
		if err != nil {
			this.AjaxReturn("0", "更新失败", nil)
			return
		}
	}
	this.AjaxReturn("1", "更新成功", nil)
	return
}
