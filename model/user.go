package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique;not null" json:"username"` //用户名
	Password string `gorm:"not null" json:"-"`               //密码
	NickName string `json:"nickname,omitempty"`              //昵称
}
