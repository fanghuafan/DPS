package models

import "github.com/jinzhu/gorm"

type ExportTarget struct {
	gorm.Model
	Id         int32
	Target     string
	Type       int32
	Sort       int32
	TemplateId int32
}
