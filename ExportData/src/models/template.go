package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Template struct {
	gorm.Model
	ID          int32
	taskId      int32
	Description string
	CreateTime  time.Time
	State       int32
}
