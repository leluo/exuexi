package admin

type Index struct {
	BaseController
}

func (c *Index) Index() {
	c.Data["stiename"] = c.sitename
	c.TplName = "admin/index.html"
}
