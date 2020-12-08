package service

import "ExportData/src/models"

type ITaskService interface {
	// add task record info
	Add(task *models.Task)

	// get task by taskCode
	GetTask(taskCode string) *models.Task
}
