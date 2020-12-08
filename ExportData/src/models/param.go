package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Param struct {
	gorm.Model
	ID         int32
	TemplateId int32
	Key        string
	Value      string
	CreateTime time.Time
}
