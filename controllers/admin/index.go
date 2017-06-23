package admin

//Index 后台首页结构体
type Index struct {
	BaseController
}

//Index 首页方法
func (c *Index) Index() {
	c.Data["stiename"] = c.sitename
	c.TplName = "admin/index.html"
}
