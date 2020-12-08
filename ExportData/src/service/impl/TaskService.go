package impl

import (
	"ExportData/src/common/logger"
	"ExportData/src/models"
	"ExportData/src/repository"
)

type TaskService struct {
	Repository repository.TaskRepository `inject:"inline"`
	Log        logger.ILogger            `inject:""`
}

// add task record info
func (t *TaskService) Add(task *models.Task) {
	t.Repository.Add(task)
}

// get task by taskCode
func (t *TaskService) GetTask(taskCode string) *models.Task {
	task := t.Repository.GetTask(taskCode)
	return task
}
