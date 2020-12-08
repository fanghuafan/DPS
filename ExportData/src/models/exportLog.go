package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ExportLog struct {
	gorm.Model
	Id         int32
	TaskId     int32
	CreateTime time.Time
}
