package admin

import (
	m "exuexi/models"
	"time"

	"strings"

	"strconv"

	"github.com/astaxie/beego/orm"
)

//User 关于后台用户的模型
type User struct {
	BaseController
}

//Login 后台用户登录的操作
func (c *User) Login() {
	var admin m.User
	if c.Ctx.Input.IsPost() {
		admin.Username = c.GetString("username")
		if err := admin.Login(c.GetString("username"), c.GetString("password")); err == nil {
			c.SetSession("admin", admin)
			admin.Ip = c.Ctx.Input.IP()
			admin.Logintime = time.Now()
			if err := admin.Update("Ip", "Logintime"); err != nil {
				c.Fail(201, "更新用户登录信息失败")
			}
			c.Redirect(c.URLFor("Index.Index"), 302)
		} else {
			c.Fail(201, err.Error())
		}
	}
	c.TplName = "admin/login.html"
}

//List 用户列表
func (c *User) List() {
	var (
		list     []*m.User
		page     int64
		pagesize int64
		cond     *orm.Condition
		count    int64
	)
	cond = new(orm.Condition)
	username := c.GetString("username")
	if username != "" {
		cond = cond.And("Username", username)
	}
	if page, _ = c.GetInt64("page"); page < 1 {
		page = 1
	}
	if pagesize, _ = c.GetInt64("pagesize"); pagesize < 1 {
		pagesize = 20
	}
	list, count = new(m.User).List(page, pagesize, cond)
	c.Data["list"] = list
	c.Data["count"] = count
	c.TplName = "admin/user-list.html"
}

//Edit 编辑用户信息
func (c *User) Edit() {
	var (
		user m.User
		err  error
	)
	if user.Id, err = c.GetInt64("id"); err != nil {
		c.Fail(201, "获取参数无效")
	}
	if err = user.Select(); err == orm.ErrNoRows {
		c.Fail(201, "改用户不存在")
	}
	if c.Ctx.Input.IsPost() {
		if err = c.ParseForm(&user); err != nil {
			c.Fail(201, "表单解析错误")
		}
		if err = user.Update("Username"); err != nil {
			c.Fail(201, "更新用户信息失败")
		}
		c.Success(200, "更新成功", "")
	}
	c.Data["user"] = user
	c.TplName = "admin/user-edit.html"
}

//Delete 删除用户信息
func (c *User) Delete() {
	ids := c.GetString("id")
	idarr := strings.Split(ids, ",")
	if len(idarr) == 0 {
		c.Fail(201, "请选择需要删除的信息")
	}
	for _, v := range idarr {
		user := new(m.User)
		user.Id, _ = strconv.ParseInt(v, 10, 64)
		if err := user.Select();err!=nil {
			c.Fail(201,"改信息不存在")
		}
		if err := user.Delete();err!=nil {
			c.Fail(201,"删除信息失败")
		}
	}
	c.Success(200,"删除成功","")
}
