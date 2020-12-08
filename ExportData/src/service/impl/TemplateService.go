package impl

import "ExportData/src/models"

type TemplateService struct {
	BaseService
}

func (templateService *TemplateService) Get(taskId int32) []models.Template {
	db := templateService.BaseService.GetDB()
	templates := []models.Template{}
	db.Where("task_id=?", taskId).Find(&templates)
	defer templateService.BaseService.CloseDB(db.DB())

	return templates
}

func (templateService *TemplateService) Add(template models.Template) {
	db := templateService.BaseService.GetDB()
	db.Create(template)
	defer templateService.BaseService.CloseDB(db.DB())
}
