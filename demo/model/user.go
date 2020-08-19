package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// 用户
type User struct {
	gorm.Model

	UUID     uuid.UUID `json:"uuid" gorm:"comment:'UUID'"`
	RoleName string    `json:"role_name" gorm:"-"`
	UserName string    `json:"user_name" gorm:"not null;unique;comment:'用户名'"`
	Password string    `json:"password" gorm:"not null;comment:'密码'"`
	Gender   string    `json:"gender" gorm:"comment:'性别'"`
	RealName string    `json:"real_name" gorm:"comment:'联系人'"`
	Mobile   string    `json:"mobile" gorm:"comment:'手机号'"`
	Address  string    `json:"address" gorm:"comment:'联系地址'"`
	IP       string    `json:"ip" gorm:"comment:'最后登录IP'"`
}
