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
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbUser := beego.AppConfig.String("db_user")
	dbPass := beego.AppConfig.String("db_pass")
	dbName := beego.AppConfig.String("db_name")
	dbTimeZone := beego.AppConfig.String("db_timezone")
	dbTablePrefix := beego.AppConfig.String("db_table_prefix")
	if dbPort == "" {
		dbPort = "3306"
	}
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if dbTimeZone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(dbTimeZone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	fmt.Println(dsn)
	orm.RegisterModelWithPrefix(dbTablePrefix,
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
