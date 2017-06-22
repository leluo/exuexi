package models

import (
	"fmt"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	//db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_timezone := beego.AppConfig.String("db_timezone")
	db_table_prefix := beego.AppConfig.String("db_table_prefix")
	if db_port == "" {
		db_port = "3306"
	}
	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8"
	if db_timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(db_timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	fmt.Println(dsn)
	orm.RegisterModelWithPrefix(db_table_prefix,
		new(User),
	)
	if beego.AppConfig.String("runmode") == "shanxi" {
		orm.Debug = true
	}
	name := "default"
	// drop table 后再建表
	force := true
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
