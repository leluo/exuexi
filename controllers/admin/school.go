package admin

//School 培训机构结构体
type School struct {
	BaseController
}

//Add 添加培训机构
func (c *School) Add() {
	c.TplName = "admin/add-school.html"
}

//Edit 编辑机构信息
func (c *BaseController) Edit() {
	c.TplName = "admin/school-edit.html"
}

//List 培训机构列表
func (c *BaseController) List() {
	c.TplName = "admin/school-list.html"
}
