package impl

import (
	"ExportData/src/models"
	"log"
)

// 导出数据服务
type ExportService struct {
	BaseService
}

// 通过模板ID获取导出操作信息
func (exportService *ExportService) Get(templateId int32) []models.ExportTarget {
	db := exportService.BaseService.GetDB()

	list := []models.ExportTarget{}
	//调用原生sql语句
	db.Raw("SELECT * FROM t_export_target WHERE template_id=?", templateId).Find(&list)
	log.Println("query export target data by template_id: ", list)
	defer exportService.BaseService.CloseDB(db.DB())

	return list
}

// 增加模板信息
func (exportService *ExportService) Add(target *models.ExportTarget) {
	db := exportService.BaseService.GetDB()
	db.Create(&target)
	defer exportService.BaseService.CloseDB(db.DB())
}

// 新增导出数据
func (exportService *ExportService) AddLog(log *models.ExportLog) {
	db := exportService.BaseService.GetDB()
	db.Create(log)
	defer exportService.BaseService.CloseDB(db.DB())
}
