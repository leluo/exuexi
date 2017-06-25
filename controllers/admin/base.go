package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/astaxie/beego"
)

//BaseController 常用函数封装
type BaseController struct {
	beego.Controller
	sitename string
	models   string
	action   string
}

//Prepare 函数执行之前需要执行的函数
func (c *BaseController) Prepare() {
	c.sitename = beego.AppConfig.String("sitename")
	c.models, c.action = c.GetControllerAndAction()
	c.Display()
}

//Auth 登录认证
func (c *BaseController) Auth() {
	if c.models == "User" && c.action == "Login" {
		return
	}
	admin := c.GetSession("admin")
	if admin == nil {
		c.Redirect(c.URLFor("User.Login", "redirect", url.QueryEscape(c.Ctx.Input.URI())), 302)
	}
}

//Success 成功时操作的函数
func (c *BaseController) Success(code int64, msg, url string) {
	if c.IsAjax() {
		c.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
			"url":  url,
		}
		c.ServeJSON()
	} else {
		c.Data["title"] = "操作成功"
		c.Data["msg"] = msg
		c.Data["code"] = code
		c.Data["url"] = url
		c.TplName = "admin/success.html"
	}
	c.StopRun()

}

//Fail 操作失败的时候调用
func (c *BaseController) Fail(code int64, msg string) {
	if c.IsAjax() {
		c.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
		}
		c.ServeJSON()
	} else {
		c.Data["title"] = "操作失败"
		c.Data["msg"] = msg
		c.TplName = "admin/fail.html"
	}
	c.StopRun()
}

//Resp 返回json格式
func (c *BaseController) Resp(code int64, msg string, data interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	c.ServeJSON()
	c.StopRun()
}

//GetMenu 获取菜单列表
func (c *BaseController) GetMenu() (list []*Menu) {
	data, err := ioutil.ReadFile("conf/menu.json")
	if err != nil {
		return list
	}
	if err := json.Unmarshal(data, &list); err != nil {
		return list
	}
	return list
}

//Display 模板公共部分
func (c *BaseController) Display() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header"] = "admin/layout/sections/header.html"
	c.LayoutSections["footer"] = "admin/layout/sections/footer.html"
	c.Data["menus"] = c.GetMenu()
	c.Layout = "admin/layout/layout.html"
}
