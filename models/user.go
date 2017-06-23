package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//User 用户表结构体
type User struct {
	Id        int64     `orm:"auto;pk"`
	Username  string    `orm:"size(60)"`
	Password  string    `orm:"size(32)"`
	Logintime time.Time `orm:"auto_now;type(date)"`
	Level     int64
	Ip        string `orm:"size(16)"`
}

//Login 用户登录
func (m *User) Login(username, password string) error {
	if err := orm.NewOrm().Read(m, "Username"); err == orm.ErrNoRows {
		return errors.New("改用户不存在")
	}
	if m.Password != password {
		return errors.New("密码不对")
	}
	return nil
}

//List 获取用户列表
func (m *User) List(page, pagesize int64, cond *orm.Condition) (list []*User, count int64) {
	if page < 1 {
		page = 1
	}
	if pagesize < 1 {
		pagesize = 20
	}
	offset := (page - 1) * pagesize
	orm.NewOrm().QueryTable(m).RelatedSel().SetCond(cond).Limit(pagesize, offset).All(&list)
	count, _ = orm.NewOrm().QueryTable(m).SetCond(cond).Count()
	return list, count
}

//Add 添加用户信息
func (m *User) Add() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

//Select 查询用户信息
func (m *User) Select(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

//Update 更新用户信息
func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

//Delete 删除用户信息
func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

//Query 使用queryseter进行查询用户信息
func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
