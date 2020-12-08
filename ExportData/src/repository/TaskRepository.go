package repository

import (
	"ExportData/src/common/logger"
	"ExportData/src/models"
)

//UserRepository 注入IDb
type TaskRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

//Add 增加人物
func (a *TaskRepository) Add(model interface{}) bool {
	if err := a.Base.Create(model); err != nil {
		a.Log.Errorf("Add task error:", err)
		return false
	}
	return true
}

//GetTask 查询人物
func (a *TaskRepository) GetTask(where interface{}) *models.Task {
	var task models.Task
	if err := a.Base.First(where, &task); err != nil {
		a.Log.Errorf("Query task error:", err)
	}
	return &task
}
