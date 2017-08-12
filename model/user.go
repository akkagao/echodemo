package model

import (
	"time"
)

type User struct {
	Id         int64      `db:"id"`
	Username   string     `db:"username"`   //昵称
	Account    string     `db:"account"`      //登录账号
	Password   string     `db:"password"`   //密码
	PhoneNum   string     `db:"phone_num"`  //电话号
	Status     int        `db:"status"`     //状态 1：系统初始化 2：用户已经完善信息 3：封号
	CreateTime *time.Time `db:"createTime"` // 注册时间
}
