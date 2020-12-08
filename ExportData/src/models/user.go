package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id         int32     `gorm:"primary_key" json:"id"`
	Name       string    `json:"username"`
	Nickname   string    `json:"nickname"`
	Password   string    `json:"password"`
	Role       int       `json:"role"`
	CreateTime time.Time `json:"create_time"`
	State      int       `json:"state"`
	Avatar     string    `json:"avatar"`
}

//BeforeCreate CreatedOn赋值
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateTime", time.Now())
	return nil
}

func (user *User) TableName() string {
	return "t_user"
}