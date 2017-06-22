package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int64     `orm:"auto;pk"`
	Username  string    `orm:"size(60)"`
	Password  string    `orm:"size(32)"`
	Logintime time.Time `orm:"auto_now;type(date)"`
	Level     int64
	Ip        string `orm:"size(16)"`
}

func (m *User) Login(username, password string) error {
	if err := orm.NewOrm().Read(m, "Username"); err == orm.ErrNoRows {
		return errors.New("改用户不存在")
	}
	if m.Password != password {
		return errors.New("密码不对")
	}
	return nil
}

func (m *User) Add() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Select(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
