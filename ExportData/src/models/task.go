package models

import (
	"time"
)

// 任务编码
type Task struct {
	Id          int32     `gorm:"primary_key" json:"id"`
	TaskCode    string    `json:"task_code"`
	TaskName    string    `json:"name"`
	Description string    `json:"desc"`
	CreateTime  time.Time `json:"create_time"`
	Creator     int32     `json:"user"`
	State       int32     `json:"state"`
	Timeout     int32     `json:"timeout"`
}

func (t *Task) TableName() string {
	return "t_task"
}