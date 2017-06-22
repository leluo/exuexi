package routers

import (
	"exuexi/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoPrefix("/admin", &admin.User{})
	beego.AutoPrefix("/admin", &admin.Index{})
}
