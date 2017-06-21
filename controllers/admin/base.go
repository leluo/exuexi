package admin

import (
	"github.com/astaxie/beego"
)

//BaseController 常用函数封装
type BaseController struct {
	beego.Controller
}

//Prepare 函数执行之前需要执行的函数
func (c *BaseController) Prepare() {

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
