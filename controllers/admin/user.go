package admin

import (
	m "exuexi/models"
)

//User 关于后台用户的模型
type User struct {
	BaseController
}

//Login 后台用户登录的操作
func (c *User) Login() {
	if c.Ctx.Input.IsPost() {
		var admin m.User
		admin.Username = c.GetString("username")
		if err := admin.Login(c.GetString("username"), c.GetString("password")); err == nil {
			c.SetSession("admin", admin)
		} else {
			c.Fail(201, err.Error())
		}
	}
	c.TplName = "admin/login.html"
}

//List 用户列表
func (c *User) List() {
	c.TplName = "admin/user-list.html"
}

//Edit 编辑用户信息
func (c *User) Edit() {
	c.TplName = "admin/user-edit.html"
}
